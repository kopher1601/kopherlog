package jp.co.kopher.kopherlog.repository

import jp.co.kopher.kopherlog.domain.User
import org.springframework.data.repository.CrudRepository

interface UserRepository : CrudRepository<User, Long> {
    fun findByEmailAndPassword(email: String, password: String): User?
}