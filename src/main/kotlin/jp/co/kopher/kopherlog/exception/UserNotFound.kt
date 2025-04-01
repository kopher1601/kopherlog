package jp.co.kopher.kopherlog.exception

class UserNotFound : KopherlogException("User not found") {
    override val code: Int
        get() = 400
}