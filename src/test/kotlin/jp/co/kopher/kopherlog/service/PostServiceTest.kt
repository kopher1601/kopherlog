package jp.co.kopher.kopherlog.service

import jp.co.kopher.kopherlog.domain.Post
import jp.co.kopher.kopherlog.exception.PostNotFound
import jp.co.kopher.kopherlog.repository.PostRepository
import jp.co.kopher.kopherlog.request.PostCreate
import jp.co.kopher.kopherlog.request.PostEdit
import jp.co.kopher.kopherlog.request.PostSearch
import org.assertj.core.api.Assertions.assertThat
import org.assertj.core.api.Assertions.assertThatThrownBy
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
        val requestPosts = (0 until 20).map {
            Post(
                _title = "吉祥寺 $it",
                _content = "マンション購入 $it",
            )
        }
        postRepository.saveAll(requestPosts)
        val search = PostSearch(
//            page = 1,
//            size = 10,
        )

        // when
        val response = postService.getList(search)

        // then
        assertThat(response.size).isEqualTo(10)
        assertThat(response.get(0).title).isEqualTo("吉祥寺 19")
        assertThat(response.get(4).title).isEqualTo("吉祥寺 15")
    }

    @Test
    @DisplayName("글 제목 수정")
    fun test4() {
        // given
        val post = Post(
            _title = "123456789012345",
            _content = "bar",
        )
        postRepository.save(post)

        val postEdit = PostEdit(
            title = "武蔵境マンションもいいな",
            content = "bar",
        )

        // when
        postService.edit(post.id!!, postEdit)

        // then
        val updatedPost = postRepository.findById(post.id!!).get()
        assertThat(updatedPost.title).isEqualTo("武蔵境マンションもいいな")
        assertThat(updatedPost.content).isEqualTo("bar")
    }

    @Test
    @DisplayName("게시글 삭제")
    fun test6() {
        // given
        val post = Post(
            _title = "123456789012345",
            _content = "bar",
        )
        postRepository.save(post)

        // when
        postService.delete(post.id!!)

        // then
        assertThat(postRepository.findById(post.id!!)).isNotPresent
    }

    @Test
    @DisplayName("getPost")
    fun test7() {
        // given
        val post = Post(
            _title = "123456789012345",
            _content = "bar",
        )
        postRepository.save(post)

        // expected
        assertThatThrownBy { postService.get(2L) }
            .isInstanceOf(PostNotFound::class.java)
            .hasMessage("Post not found")
    }

}