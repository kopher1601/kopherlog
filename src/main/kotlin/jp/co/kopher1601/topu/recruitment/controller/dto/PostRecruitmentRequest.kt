package jp.co.kopher1601.topu.recruitment.controller.dto

import jp.co.kopher1601.topu.recruitment.domain.ProgressMethods
import jp.co.kopher1601.topu.recruitment.domain.RecruitmentCategories
import java.time.LocalDate

data class PostRecruitmentRequest(
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
)