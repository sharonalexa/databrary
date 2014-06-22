package models

import scala.concurrent.Future
import play.api.libs.concurrent.Execution.Implicits.defaultContext
import macros._
import dbrary._
import site._

/** Link between Authorize (rows in table/view) and Access. */
sealed class Authorization (val child : Party, val parent : Party, val site : Permission.Value, val member : Permission.Value) extends Access with TableRow {
  final def childId = child.id
  final def parentId = parent.id
  private[models] def sqlKey = SQLTerms('child -> childId, 'parent -> parentId)

  final def target = parent
  final def identity = child
}

/** A specific authorization by one party of another.
  * An authorization represents granting of certain permissions to a party by/under another, which can also be viewed as group membership.
  * Unlike most [TableRow] classes, these may be constructed directly, and thus are not expected to reflect the current state of the database in the same way and have different update semantics.
  * @constructor create an authorization object, not (yet) persisted to the database
  * @param child the party being authorized, the "member"
  * @param parent the party whose permissions are being authorized, the "group"
  * @param site the level of site/group access granted via the parent to the child, thus the maximum site permission level inherited by the child from the parent
  * @param member the specific permissions granted on behalf of the parent to the child, such that the child has rights to perform actions up to this permission as the parent (not inherited)
  * @param expires the time at which this authorization stops, or never if None
  */
final class Authorize protected (child : Party, parent : Party, site : Permission.Value, member : Permission.Value, val expires : Option[Timestamp], val info : Option[String]) extends Authorization(child, parent, site, member) {
  /** Determine if this authorization is currently in effect.
    * @return true if expires is unset or in the future */
  def valid = expires.forall(_.toDateTime.isAfterNow)
  /** True if this authorization has been "granted." */
  def authorized = site != Permission.NONE || member != Permission.NONE

  def json = JsonObject.flatten(
    Some('site -> site),
    Some('member -> member),
    expires.map('expires -> _),
    info.map('info -> _)
  )
}

object Authorize extends Table[Authorize]("authorize") {
  private val columns = Columns(
      SelectColumn[Permission.Value]("site")
    , SelectColumn[Permission.Value]("member")
    , SelectColumn[Option[Timestamp]]("expires")
    ).~?(Info.columns)
    .map { case ((site, member, expires), info) =>
      (child : Party, parent : Party) => new Authorize(child, parent, site, member, expires, info)
    }

  private[this] val condition = "AND (expires IS NULL OR expires > CURRENT_TIMESTAMP)"
  private[this] def conditionIf(all : Boolean) =
    if (all) "" else condition

  def get(child : Party, parent : Party) : Future[Option[Authorize]] =
    columns
      .map(_(child, parent))
      .SELECT("WHERE child = ? AND parent = ?")
      .apply(child.id, parent.id).singleOpt

  /** Get all authorizations granted to a particular child.
    * @param all include inactive authorizations
    */
  private[models] def getParents(child : Party, all : Boolean = false) : Future[Seq[Authorize]] =
    columns.join(Party.row, "parent = party.id")
      .map { case (a, p) => a(child, p) }
      .SELECT("WHERE child = ?", conditionIf(all)).apply(child.id).list
  /** Get all authorizations granted ba a particular parent.
    * @param all include inactive authorizations
    */
  private[models] def getChildren(parent : Party, all : Boolean = false) : Future[Seq[Authorize]] =
    columns.join(Party.row, "child = party.id")
      .map { case (a, c) => a(c, parent) }
      .SELECT("WHERE parent = ?", conditionIf(all)).apply(parent.id).list

  def getAll() : Future[Seq[Authorize]] =
    columns
    .join(Party.columns.fromAlias("child"), "child = child.id")
    .join(Party.columns.fromAlias("parent"), "parent = parent.id")
    .map { case ((auth, child), parent) => auth(child, parent) }
    .SELECT("ORDER BY parent.id, site")
    .apply().list

  /** Update or add a specific authorization in the database.
    * If an authorization for the child and parent already exist, it is changed to match this.
    * Otherwise, a new one is added.
    * This may invalidate child.access. */
  def set(child : Party.Id, parent : Party.Id, site : Permission.Value, member : Permission.Value, expires : Option[Timestamp] = None)(implicit request : Site) : Future[Boolean] =
    Audit.changeOrAdd(Authorize.table, SQLTerms('site -> site, 'member -> member, 'expires -> expires), SQLTerms('child -> child, 'parent -> parent)).execute
  /** Remove a particular authorization from the database.
    * @return true if a matching authorization was found and deleted
    */
  def delete(child : Party.Id, parent : Party.Id)(implicit site : Site) : Future[Boolean] =
    Audit.remove("authorize", SQLTerms('child -> child, 'parent -> parent)).execute

  object Info extends Table[String]("authorize_info") {
    private[Authorize] val columns = Columns(
      SelectColumn[String]("info")
    )
    def set(child : Party.Id, parent : Party.Id, info : Option[String]) : Future[Boolean] = {
      val id = SQLTerms('child -> child, 'parent -> parent)
      info.fold(DELETE(id)) { info =>
	DBUtil.updateOrInsert(
	  SQL("UPDATE", table, "SET info = ? WHERE", id.where)(_, _).apply(info +: id))(
	  INSERT(id :+ ('info -> info))(_, _))
      }.execute
    }
  }
}

object Authorization extends Table[Authorization]("authorize_view") {
  private[models] val columns = Columns(
      SelectColumn[Permission.Value]("site")
    , SelectColumn[Permission.Value]("member")
    )

  private def unOpt(access : Option[(Permission.Value, Permission.Value)]) : (Permission.Value, Permission.Value) =
    access.getOrElse((Permission.NONE, Permission.NONE))

  object Nobody extends Authorization(Party.Nobody, Party.Root, Permission.NONE, Permission.NONE)
  object Root extends Authorization(Party.Root, Party.Root, Permission.ADMIN, Permission.ADMIN)

  private final class Self (party : Party) extends Authorization(party, party,
    if (party.id === Party.NOBODY) Permission.NONE else Permission.ADMIN,
    if (party.id === Party.NOBODY) Permission.NONE else Permission.ADMIN)

  private[models] def make(child : Party, parent : Party = Party.Root)(access : Option[(Permission.Value, Permission.Value)]) : Authorization = {
    val (site, member) = unOpt(access)
    new Authorization(child, parent, site, member)
  }

  /** Determine the effective inherited and direct permission levels granted to a child by a parent. */
  private[models] def get(child : Party, parent : Party = Party.Root) : Future[Authorization] =
    if (child === parent) async(new Self(parent)) // optimization
    else columns
      .SELECT("WHERE child = ? AND parent = ?")
      .apply(child.id, parent.id).singleOpt
      .map(make(child, parent))

  def _get(childId : Party.Id, parent : Party = Party.Root) : Future[Option[Authorization]] =
    if (childId === parent.id) async(Some(new Self(parent))) // optimization
    else if (childId == Party.NOBODY) async(Some(Nobody))
    else if (childId == Party.ROOT) async(Some(Root))
    else Party.row
      .leftJoin(columns, "party.id = authorize_view.child AND authorize_view.parent = ?")
      .map { case (p, a) => make(p)(a) }
      .SELECT("WHERE party.id = ?")
      .apply(parent.id, childId).singleOpt
}
