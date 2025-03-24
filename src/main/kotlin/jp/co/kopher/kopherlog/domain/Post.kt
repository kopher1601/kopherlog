package jp.co.kopher.kopherlog.domain

import jakarta.persistence.*

@Entity
class Post(

    private var title: String,

    @Lob
    private var content: String,

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private var id: Long? = null,
)