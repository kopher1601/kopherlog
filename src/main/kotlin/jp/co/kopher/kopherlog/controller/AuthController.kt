package jp.co.kopher.kopherlog.controller

import jp.co.kopher.kopherlog.service.AuthService
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class AuthController(
    private val authService: AuthService,
) {

    val log: Logger = LoggerFactory.getLogger(this.javaClass)

    @GetMapping("/auth/login")
    fun login(): String {
        return "login page"
    }

}