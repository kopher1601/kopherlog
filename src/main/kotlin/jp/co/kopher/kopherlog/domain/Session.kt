package jp.co.kopher.kopherlog.domain

import jakarta.persistence.*
import java.util.*

@Entity
class Session(

    private val _accessToken: String = UUID.randomUUID().toString(),

    @ManyToOne
    private val user: User,

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private val id: Long? = null,
) {

    val accessToken: String
        get() = _accessToken
}