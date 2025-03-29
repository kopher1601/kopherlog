package jp.co.kopher.kopherlog.exception

class InvalidRequest(
    fieldName: String,
    errorMessage: String,
) : KopherlogException("Invalid request") {
    override val code: Int
        get() = 400

    init {
        super.addValidation(fieldName, errorMessage)
    }
}