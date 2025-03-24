package jp.co.kopher.kopherlog.response

/**
 * {
 *      "code": "400",
 *      "message": "必須です。"
 *      "validation" : {
 *          "title": "タイトルを入力してください。"
 *      }
 * }
 */
data class ErrorResponse(
    private var _code: String,
    private var _message: String,
    private val _validation: MutableMap<String, String> = mutableMapOf()
) {
    fun addValidation(fieldName: String, errorMessage: String) {
        this._validation[fieldName] = errorMessage
    }

    val code: String
        get() = _code

    val message: String
        get() = _message

    val validation: Map<String, String>
        get() = _validation
}