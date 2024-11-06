package jp.co.kopher1601.topu.recruitment.service.dto

import jakarta.validation.constraints.NotNull
import jp.co.kopher1601.topu.recruitment.domain.ProgressMethods
import jp.co.kopher1601.topu.recruitment.domain.RecruitmentCategories
import jp.co.kopher1601.topu.recruitment.controller.dto.PostRecruitmentRequest
import java.time.LocalDate

data class PostRecruitment(
    @field:NotNull
    val recruitmentCategories: RecruitmentCategories?,
    val progressMethods: ProgressMethods?,
    val techStacks: List<String>?,
    val recruitmentPositions: List<String>?,
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
                request.recruitmentCategories,
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
