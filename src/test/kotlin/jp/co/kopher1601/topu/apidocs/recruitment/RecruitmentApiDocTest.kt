package jp.co.kopher1601.topu.apidocs.recruitment

import com.fasterxml.jackson.databind.ObjectMapper
import jp.co.kopher1601.topu.recruitment.domain.ProgressMethods
import jp.co.kopher1601.topu.recruitment.domain.RecruitmentCategories
import jp.co.kopher1601.topu.recruitment.request.PostRecruitmentRequest
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.restdocs.AutoConfigureRestDocs
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.http.MediaType.APPLICATION_JSON
import org.springframework.restdocs.mockmvc.MockMvcRestDocumentation.document
import org.springframework.restdocs.mockmvc.RestDocumentationRequestBuilders.post
import org.springframework.restdocs.operation.preprocess.Preprocessors.*
import org.springframework.restdocs.payload.PayloadDocumentation.fieldWithPath
import org.springframework.restdocs.payload.PayloadDocumentation.requestFields
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.result.MockMvcResultHandlers.print
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.status
import java.time.LocalDate

@SpringBootTest
@AutoConfigureRestDocs
@AutoConfigureMockMvc
class RecruitmentApiDocTest @Autowired constructor(
    val mvc: MockMvc,
    val objectMapper: ObjectMapper,
) {

    @Test
    @DisplayName("응모글을 작성하면 응모글 목록에 담긴다")
    fun postRecruitment() {
        // given
        val request = PostRecruitmentRequest(
            RecruitmentCategories.STUDY,
            ProgressMethods.ALL,
            listOf("Kotlin", "Spring", "JPA"),
            listOf("バックエンド", "SRE", "アプリケーション開発者"),
            3,
            6,
            LocalDate.of(2024, 11, 4),
            "test@test.com",
            "Kotlin勉強会を一緒にしましょう",
            "Kotlin 基礎本からスタートします。",
        )
        val jsonString = objectMapper.writeValueAsString(request)

        // expected
        mvc.perform(
            post("/recruitments")
                .contentType(APPLICATION_JSON)
                .content(jsonString)
        )
            .andExpect(status().isCreated())
            .andDo(
                document(
                    "recruitment-post",
                    preprocessRequest(prettyPrint()),
                    preprocessResponse(prettyPrint()),
                    requestFields(
                        fieldWithPath("recruitmentCategories").description("응모 카테고리"),
                        fieldWithPath("progressMethods").description("응모 방법"),
                        fieldWithPath("techStacks").description("기술 스택"),
                        fieldWithPath("recruitmentPositions").description("응모 포지션"),
                        fieldWithPath("numberOfPeople").description("모집 인원"),
                        fieldWithPath("progressPeriod").description("진행 기간"),
                        fieldWithPath("recruitmentDeadline").description("응모 마감일"),
                        fieldWithPath("contract").description("연락처"),
                        fieldWithPath("subject").description("제목"),
                        fieldWithPath("content").description("내용")
                    )
                )
            )
            .andDo(print())
    }
}