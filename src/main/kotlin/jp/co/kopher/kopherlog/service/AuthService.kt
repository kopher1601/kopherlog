package jp.co.kopher.kopherlog.service

import jp.co.kopher.kopherlog.repository.UserRepository
import org.springframework.stereotype.Service

@Service
class AuthService(
    private val userRepository: UserRepository,
)