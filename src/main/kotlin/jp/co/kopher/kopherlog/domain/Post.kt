package jp.co.kopher.kopherlog.domain

import jakarta.persistence.*
import jp.co.kopher.kopherlog.request.PostEdit

@Entity
class Post(

    @Column(name = "title")
    private var _title: String,

    @Lob
    @Column(name = "content")
    private var _content: String,

    @Id
    @Column(name = "id")
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private val _id: Long? = null,
) {
    fun edit(postEdit: PostEdit) {
        _title = postEdit.title
        _content = postEdit.content
    }

    val title: String
        get() = _title

    val content: String
        get() = _content

    val id: Long?
        get() = _id
}