package jp.co.kopher.kopherlog.request

import jakarta.validation.constraints.NotBlank

data class PostEdit(
    @field:NotBlank(message = "タイトルを入力してください。")
    var title: String,

    @field:NotBlank(message = "コンテンツを入力してください。")
    var content: String,
)