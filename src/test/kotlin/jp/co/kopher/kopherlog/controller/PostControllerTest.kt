package jp.co.kopher.kopherlog.controller

import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest
import org.springframework.test.context.TestConstructor
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders
import org.springframework.test.web.servlet.result.MockMvcResultHandlers.print
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.content
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.status

@WebMvcTest
@TestConstructor(autowireMode = TestConstructor.AutowireMode.ALL)
class PostControllerTest(
    private val mockMvc: MockMvc,
) {

    @Test
    @DisplayName("/posts 요청시 Hello world 를 출력한다")
    fun test() {
        // expected
        mockMvc.perform(MockMvcRequestBuilders.get("/posts"))
            .andExpect(status().isOk)
            .andExpect(content().string("Hello World"))
            .andDo(print())
    }

}