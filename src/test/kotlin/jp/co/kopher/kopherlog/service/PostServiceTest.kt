package jp.co.kopher.kopherlog.service

import jp.co.kopher.kopherlog.domain.Post
import jp.co.kopher.kopherlog.repository.PostRepository
import jp.co.kopher.kopherlog.request.PostCreate
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.test.context.TestConstructor
import org.springframework.transaction.annotation.Transactional

@SpringBootTest
@Transactional
@TestConstructor(autowireMode = TestConstructor.AutowireMode.ALL)
class PostServiceTest(
    private val postService: PostService,
    private val postRepository: PostRepository,
) {

    @Test
    fun test1() {
        // given
        val postCreate = PostCreate(
            title = "武蔵境マンションもいいな",
            content = "武蔵境マンション購入"
        )

        // when
        postService.write(postCreate)

        // then
        assertThat(postRepository.count()).isEqualTo(1)
        val posts = postRepository.findAll()
        assertThat(posts[0].title).isEqualTo(postCreate.title)
        assertThat(posts[0].content).isEqualTo(postCreate.content)
    }

    @Test
    fun test2() {
        // given
        val post = Post(
            _title = "123456789012345",
            _content = "bar",
        )
        postRepository.save(post)

        // when
        val response = postService.get(post.id!!)

        // then
        assertThat(response.title).isEqualTo("1234567890")
        assertThat(response.content).isEqualTo(post.content)
    }

    @Test
    @DisplayName("글 여러 개 조회")
    fun test3() {
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
        val response = postService.getList()

        // then
        assertThat(response.size).isEqualTo(2)
    }
}