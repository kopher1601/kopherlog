package jp.co.kopher.kopherlog.controller

import com.fasterxml.jackson.databind.ObjectMapper
import jp.co.kopher.kopherlog.request.PostCreate
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest
import org.springframework.http.MediaType
import org.springframework.test.context.TestConstructor
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post
import org.springframework.test.web.servlet.result.MockMvcResultHandlers.print
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.*

@WebMvcTest
@TestConstructor(autowireMode = TestConstructor.AutowireMode.ALL)
class PostControllerTest(
    private val mockMvc: MockMvc,
    private val objectMapper: ObjectMapper,
) {

    @Test
    @DisplayName("/posts 요청시 Hello world 를 출력한다")
    fun test() {
        // given
        val request = PostCreate("吉祥寺マンション", "コンテンツ")
        val jsonString = objectMapper.writeValueAsString(request)

        // expected
        mockMvc.perform(
            post("/posts")
                .contentType(MediaType.APPLICATION_JSON)
                .content(jsonString)
        ).andExpect(status().isOk)
            .andExpect(content().string("Hello World"))
            .andDo(print())
    }

    @Test
    @DisplayName("/posts 요청시 title 값은 필수이다")
    fun test2() {
        // given
        val request = "{\"title\":  \"\", \"content\": \"内容\"}"

        // expected
        mockMvc.perform(
            post("/posts")
                .contentType(MediaType.APPLICATION_JSON)
                .content(request)
        )
            .andExpect(status().isBadRequest)
            .andExpect(jsonPath("$.code").value("400"))
            .andExpect(jsonPath("$.message").value("リクエストに誤りがあります。"))
            .andExpect(jsonPath("$.validation.title").value("タイトルを入力してください。"))
            .andDo(print())
    }

}