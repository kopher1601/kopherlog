package jp.co.kopher.kopherlog.exception

abstract class KopherlogException(
    message: String,
    val validation: MutableMap<String, String> = mutableMapOf(),
) : RuntimeException(message) {

    abstract val code: Int

    fun addValidation(fieldName: String, errorMessage: String) {
        validation[fieldName] = errorMessage
    }

}