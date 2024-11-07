package jp.co.kopher1601.topu.recruitment.service.dto

import jakarta.validation.constraints.NotNull
import jp.co.kopher1601.topu.recruitment.domain.ProgressMethods
import jp.co.kopher1601.topu.recruitment.domain.RecruitmentCategory
import jp.co.kopher1601.topu.recruitment.controller.dto.PostRecruitmentRequest
import java.time.LocalDate

data class PostRecruitment(
    @field:NotNull
    val recruitmentCategory: RecruitmentCategory?,
    val progressMethods: ProgressMethods?,
    val techStacks: List<String>?,
    val positions: List<String>?,
    val numberOfPeople: Int?,
    val progressPeriod: Int?,
    val recruitmentDeadline: LocalDate?,
    val contract: String?,
    val subject: String?,
    val content: String?,
) {

    companion object {
        fun from(request: PostRecruitmentRequest): PostRecruitment {
            return PostRecruitment(
                request.recruitmentCategory,
                request.progressMethods,
                request.techStacks,
                request.recruitmentPositions,
                request.numberOfPeople,
                request.progressPeriod,
                request.recruitmentDeadline,
                request.contract,
                request.subject,
                request.content,
            )
        }
    }
}
