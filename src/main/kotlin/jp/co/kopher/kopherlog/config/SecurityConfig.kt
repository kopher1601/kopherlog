package jp.co.kopher.kopherlog.config

import org.springframework.boot.autoconfigure.security.servlet.PathRequest.toH2Console
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.HttpMethod
import org.springframework.security.config.annotation.web.builders.HttpSecurity
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity
import org.springframework.security.config.annotation.web.configuration.WebSecurityCustomizer
import org.springframework.security.web.SecurityFilterChain

@Configuration
@EnableWebSecurity
class SecurityConfig {

    @Bean
    fun webSecurityCustomizer(): WebSecurityCustomizer {
        return WebSecurityCustomizer {
            it.ignoring()
                .requestMatchers("/favicon.ico")
                .requestMatchers("/error")
                .requestMatchers(toH2Console())
        }
    }

    @Bean
    fun securityFilterChain(http: HttpSecurity): SecurityFilterChain {
        return http
            .authorizeHttpRequests {
                it.requestMatchers(HttpMethod.GET, "/auth/login").permitAll()
            }
            .csrf {
                it.disable()
            }
            .build()
    }
}