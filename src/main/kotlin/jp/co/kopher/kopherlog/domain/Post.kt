package jp.co.kopher.kopherlog.domain

import jakarta.persistence.*

@Entity
class Post(

    @Column(name = "title")
    private val _title: String,

    @Lob
    @Column(name = "content")
    private val _content: String,

    @Id
    @Column(name = "id")
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