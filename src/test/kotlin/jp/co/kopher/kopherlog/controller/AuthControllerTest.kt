package jp.co.kopher.kopherlog.controller

import com.fasterxml.jackson.databind.ObjectMapper
import jp.co.kopher.kopherlog.domain.User
import jp.co.kopher.kopherlog.exception.UserNotFound
import jp.co.kopher.kopherlog.repository.UserRepository
import jp.co.kopher.kopherlog.request.Login
import org.assertj.core.api.Assertions.assertThat
import org.hamcrest.Matchers.notNullValue
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.data.repository.findByIdOrNull
import org.springframework.http.MediaType
import org.springframework.test.context.TestConstructor
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post
import org.springframework.test.web.servlet.result.MockMvcResultHandlers.print
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.status
import org.springframework.transaction.annotation.Transactional

@SpringBootTest
@AutoConfigureMockMvc
@Transactional
@TestConstructor(autowireMode = TestConstructor.AutowireMode.ALL)
class AuthControllerTest(
    private val mockMvc: MockMvc,
    private val objectMapper: ObjectMapper,
    private val userRepository: UserRepository,
) {

    @DisplayName("로그인 성공 후 세션 1개 생성")
    @Test
    fun session() {
        // given
        val savedUser = userRepository.save(
            User(
                name = "kopher",
                email = "kopherlog@example.com",
                password = "1234",
            )
        )

        val request = Login("kopherlog@example.com", "1234")
        val jsonString = objectMapper.writeValueAsString(request)

        // expected
        mockMvc.perform(
            post("/auth/login")
                .contentType(MediaType.APPLICATION_JSON)
                .content(jsonString)
        )
            .andExpect(status().isOk)
            .andDo(print())
        val loggedInUser = userRepository.findByIdOrNull(savedUser.id!!) ?: throw UserNotFound()
        assertThat(loggedInUser.sessions.size).isEqualTo(1)
    }

    @DisplayName("로그인 성공 후 세션 응답")
    @Test
    fun sessionResponse() {
        // given
        val savedUser = userRepository.save(
            User(
                name = "kopher",
                email = "kopherlog@example.com",
                password = "1234",
            )
        )

        val request = Login("kopherlog@example.com", "1234")
        val jsonString = objectMapper.writeValueAsString(request)
        val loggedInUser = userRepository.findByIdOrNull(savedUser.id!!) ?: throw UserNotFound()

        // expected
        mockMvc.perform(
            post("/auth/login")
                .contentType(MediaType.APPLICATION_JSON)
                .content(jsonString)
        )
            .andExpect(status().isOk)
            .andExpect(jsonPath("$.accessToken", notNullValue()))
            .andExpect(jsonPath("$.accessToken").value(loggedInUser.sessions[0].accessToken))
            .andDo(print())
    }

}