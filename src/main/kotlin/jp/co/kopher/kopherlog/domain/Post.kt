package jp.co.kopher.kopherlog.domain

import jakarta.persistence.*

@Entity
class Post(

    private val _title: String,

    @Lob
    private val _content: String,

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private val _id: Long? = null,
) {
    val title: String
        get() = _title

    val content: String
        get() = _content

    val id: Long?
        get() = _id
}