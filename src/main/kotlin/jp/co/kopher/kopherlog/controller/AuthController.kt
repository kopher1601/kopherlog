package jp.co.kopher.kopherlog.controller

import jp.co.kopher.kopherlog.request.Login
import jp.co.kopher.kopherlog.response.SessionResponse
import jp.co.kopher.kopherlog.service.AuthService
import org.slf4j.LoggerFactory
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RestController

@RestController
class AuthController(
    private val authService: AuthService,
) {

    val log = LoggerFactory.getLogger(this.javaClass)

    @PostMapping("/auth/login")
    fun login(@RequestBody request: Login): SessionResponse {
        val accessToken = authService.signIn(request)
        return SessionResponse(accessToken)
    }
}