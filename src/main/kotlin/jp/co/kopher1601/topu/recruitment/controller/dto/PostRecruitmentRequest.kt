package jp.co.kopher1601.topu.recruitment.controller.dto

import jp.co.kopher1601.topu.recruitment.domain.enums.ProgressMethods
import jp.co.kopher1601.topu.recruitment.domain.enums.RecruitmentCategory
import java.time.LocalDate

data class PostRecruitmentRequest(
    val recruitmentCategory: RecruitmentCategory?,
    val progressMethods: ProgressMethods?,
    val techStacks: List<String>?,
    val recruitmentPositions: List<String>?,
    val numberOfPeople: Int?,
    val progressPeriod: Int?,
    val recruitmentDeadline: LocalDate?,
    val contract: String?,
    val subject: String?,
    val content: String?,
)