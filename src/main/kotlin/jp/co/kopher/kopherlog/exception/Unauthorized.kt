package jp.co.kopher.kopherlog.exception

class Unauthorized : KopherlogException("Unauthorized") {
    override val code: Int
        get() = 401
}