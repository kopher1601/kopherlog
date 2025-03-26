package jp.co.kopher.kopherlog.controller

import com.fasterxml.jackson.databind.ObjectMapper
import jp.co.kopher.kopherlog.domain.Post
import jp.co.kopher.kopherlog.repository.PostRepository
import jp.co.kopher.kopherlog.request.PostCreate
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.http.MediaType
import org.springframework.test.context.TestConstructor
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post
import org.springframework.test.web.servlet.result.MockMvcResultHandlers.print
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.status
import org.springframework.transaction.annotation.Transactional

@SpringBootTest
@AutoConfigureMockMvc
@Transactional
@TestConstructor(autowireMode = TestConstructor.AutowireMode.ALL)
class PostControllerTest(
    private val mockMvc: MockMvc,
    private val objectMapper: ObjectMapper,
    private val postRepository: PostRepository,
) {

    @Test
    @DisplayName("/posts 요청시 title 값은 필수이다")
    fun test2() {
        // given
        val request = PostCreate(title = "", content = "内容")
        val jsonString = objectMapper.writeValueAsString(request)

        // expected
        mockMvc.perform(
            post("/posts")
                .contentType(MediaType.APPLICATION_JSON)
                .content(jsonString)
        )
            .andExpect(status().isBadRequest)
            .andExpect(jsonPath("$.code").value("400"))
            .andExpect(jsonPath("$.message").value("リクエストに誤りがあります。"))
            .andExpect(jsonPath("$.validation.title").value("タイトルを入力してください。"))
            .andDo(print())
    }

    @Test
    @DisplayName("/posts 요청시 DB에 값이 저장된다")
    fun test3() {
        // given
        val request = PostCreate(
            title = "吉祥寺マンション",
            content = "マンション購入"
        )
        val jsonString = objectMapper.writeValueAsString(request)

        // when
        mockMvc.perform(
            post("/posts")
                .contentType(MediaType.APPLICATION_JSON)
                .content(jsonString)
        )
            .andExpect(status().isCreated)
            .andDo(print())

        // then
        assertThat(postRepository.count()).isEqualTo(1)
    }

    @Test
    @DisplayName("글 1개 조회")
    fun test4() {
        // given
        val post = Post(
            _title = "123456789012345",
            _content = "bar",
        )
        postRepository.save(post)

        // when
        mockMvc.perform(
            get("/posts/{postId}", post.id!!)
                .contentType(MediaType.APPLICATION_JSON)
        )
            .andExpect(status().isOk)
            .andExpect(jsonPath("$.title").value("1234567890"))
            .andExpect(jsonPath("$.content").value("bar"))
            .andDo(print())
    }

    @Test
    @DisplayName("글 여러 개 조회")
    fun test5() {
        // given
        val post1 = Post(
            _title = "foo1",
            _content = "bar1",
        )
        val post2 = Post(
            _title = "foo2",
            _content = "bar2",
        )
        postRepository.saveAll(listOf(post1, post2))

        // when
        mockMvc.perform(
            get("/posts")
                .contentType(MediaType.APPLICATION_JSON)
        )
            .andExpect(status().isOk)
            .andExpect(jsonPath("$.length()").value(2))
            .andExpect(jsonPath("$[0].title").value("foo1"))
            .andExpect(jsonPath("$[0].content").value("bar1"))
            .andDo(print())
    }
}