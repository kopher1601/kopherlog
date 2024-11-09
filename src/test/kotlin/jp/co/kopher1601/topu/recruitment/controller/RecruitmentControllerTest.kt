package jp.co.kopher1601.topu.recruitment.controller

import com.fasterxml.jackson.databind.ObjectMapper
import jp.co.kopher1601.topu.recruitment.domain.Recruitment
import jp.co.kopher1601.topu.recruitment.domain.enums.ProgressMethods
import jp.co.kopher1601.topu.recruitment.domain.enums.RecruitmentCategory
import jp.co.kopher1601.topu.recruitment.repository.RecruitmentRepository
import org.hamcrest.Matchers.`is`
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.http.MediaType.APPLICATION_JSON
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get
import org.springframework.test.web.servlet.result.MockMvcResultHandlers.print
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.status
import java.time.LocalDate

@SpringBootTest
@AutoConfigureMockMvc
class RecruitmentControllerTest @Autowired constructor(
    private val recruitmentRepository: RecruitmentRepository,
    private val mvc: MockMvc,
    private val objectMapper: ObjectMapper,
) {


    @Test
    @DisplayName("recruitment idを用いてrecruitmentを照会できる")
    fun getRecruitment() {

        val recruitment = Recruitment(
            subject = "チーム員を募集します。",
            content = "美しいプロジェクトを作って行きましょう。",
            contract = "test@test.com",
            recruitmentDeadLine = LocalDate.of(2024, 11, 9),
            progressPeriod = 3,
            numberOfPeople = 3,
            progressMethods = ProgressMethods.ONLINE,
            recruitmentCategory = RecruitmentCategory.PROJECT,
        )
        recruitmentRepository.save(recruitment)
        val jsonString = objectMapper.writeValueAsString(recruitment)

        mvc.perform(
            get("/recruitments/{recruitmentId}", recruitment.id)
                .contentType(APPLICATION_JSON)
                .content(jsonString)
        ).andExpect(status().isOk)
            .andExpect(jsonPath("$.id", `is`(recruitment.id!!.toInt())))
            .andExpect(jsonPath("$.subject", `is`("チーム員を募集します。")))
            .andExpect(jsonPath("$.content", `is`("美しいプロジェクトを作って行きましょう。")))
            .andExpect(jsonPath("$.contract", `is`("test@test.com")))
            .andExpect(jsonPath("$.recruitmentDeadLine", `is`(LocalDate.of(2024, 11, 9).toString())))
            .andExpect(jsonPath("$.progressPeriod", `is`(3)))
            .andExpect(jsonPath("$.numberOfPeople", `is`(3)))
            .andExpect(jsonPath("$.progressMethods", `is`(ProgressMethods.ONLINE.name)))
            .andExpect(jsonPath("$.recruitmentCategory", `is`(RecruitmentCategory.PROJECT.name)))
            .andDo(print())

    }
}