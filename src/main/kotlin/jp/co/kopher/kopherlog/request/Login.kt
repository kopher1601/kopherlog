package jp.co.kopher.kopherlog.request

import jakarta.validation.constraints.NotBlank

data class Login(
    @field:NotBlank(message = "メールを入力してください。")
    val email: String,
    @field:NotBlank(message = "パスワードを入力してください。")
    val password: String
)