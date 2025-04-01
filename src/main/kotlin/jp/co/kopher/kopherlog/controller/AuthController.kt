package jp.co.kopher.kopherlog.controller

import jp.co.kopher.kopherlog.exception.UserNotFound
import jp.co.kopher.kopherlog.repository.UserRepository
import jp.co.kopher.kopherlog.request.Login
import org.slf4j.LoggerFactory
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RestController

@RestController
class AuthController(
    private val userRepository: UserRepository,
) {

    val log = LoggerFactory.getLogger(this.javaClass)

    @PostMapping("/auth/login")
    fun login(@RequestBody request: Login) {
        log.info("login >>> {}", request)

        // DB
        val user = userRepository.findByEmailAndPassword(request.email, request.password)
            ?: throw UserNotFound()

        // token
    }
}