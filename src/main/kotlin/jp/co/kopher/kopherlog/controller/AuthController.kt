package jp.co.kopher.kopherlog.controller

import jp.co.kopher.kopherlog.service.AuthService
import org.slf4j.LoggerFactory
import org.springframework.web.bind.annotation.RestController

@RestController
class AuthController(
    private val authService: AuthService,
) {

    val log = LoggerFactory.getLogger(this.javaClass)

}