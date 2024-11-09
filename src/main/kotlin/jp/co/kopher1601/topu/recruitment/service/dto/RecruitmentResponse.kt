package jp.co.kopher1601.topu.recruitment.service.dto

import jp.co.kopher1601.topu.recruitment.domain.Recruitment
import jp.co.kopher1601.topu.recruitment.domain.enums.ProgressMethods
import jp.co.kopher1601.topu.recruitment.domain.enums.RecruitmentCategory
import java.time.LocalDate

data class RecruitmentResponse(
    val positions: List<String> = listOf(),
    val techStacks: List<String> = listOf(),
    val content: String,
    val pageViews: Long,
    val subject: String,
    val contract: String,
    val recruitmentDeadLine: LocalDate,
    val progressPeriod: Int,
    val numberOfPeople: Int,
    val progressMethods: ProgressMethods,
    val recruitmentCategory: RecruitmentCategory,
    val id: Long,
) {
    companion object {
        fun from(recruitment: Recruitment): RecruitmentResponse {
            return RecruitmentResponse(
                positions = recruitment.recruitmentPositions.map { rp ->
                    rp.position.positionName}.toList(),
                techStacks = recruitment.recruitmentTechStacks.map { rt ->
                    rt.techStack.technologyName }.toList(),
                content = recruitment.content,
                pageViews = recruitment.pageViews,
                subject = recruitment.subject,
                contract = recruitment.contract,
                recruitmentDeadLine = recruitment.recruitmentDeadLine,
                progressPeriod = recruitment.progressPeriod,
                numberOfPeople = recruitment.numberOfPeople,
                progressMethods = recruitment.progressMethods,
                recruitmentCategory = recruitment.recruitmentCategory,
                id = recruitment.id!!,
            )
        }
    }
}