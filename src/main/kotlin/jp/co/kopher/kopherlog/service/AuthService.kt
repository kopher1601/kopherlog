package jp.co.kopher.kopherlog.service

import jp.co.kopher.kopherlog.exception.UserNotFound
import jp.co.kopher.kopherlog.repository.UserRepository
import jp.co.kopher.kopherlog.request.Login
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional

@Service
class AuthService(
    private val userRepository: UserRepository,
) {

    @Transactional
    fun signIn(login: Login): String {
        val user = userRepository.findByEmailAndPassword(login.email, login.password)
            ?: throw UserNotFound()
        val session = user.addSession()

        return session.accessToken
    }
}