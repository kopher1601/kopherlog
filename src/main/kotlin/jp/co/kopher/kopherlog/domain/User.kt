package jp.co.kopher.kopherlog.domain

import jakarta.persistence.*
import java.time.LocalDateTime

@Entity
@Table(name = "users")
class User(
    val name: String,
    val email: String,
    val password: String,
    val createdAt: LocalDateTime? = null,

    @Column(name = "sessions")
    @OneToMany(cascade = [CascadeType.ALL], mappedBy = "user")
    private val _sessions: MutableList<Session> = mutableListOf(),

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long? = null,
) {
    fun addSession(): Session {
        val session = Session(user = this)
        _sessions.add(session)

        return session
    }

    val sessions: List<Session>
        get() = _sessions
}