package jp.co.kopher1601.topu.recruitment.domain

import jakarta.persistence.*
import jp.co.kopher1601.topu.recruitment.domain.enums.ProgressMethods
import jp.co.kopher1601.topu.recruitment.domain.enums.RecruitmentCategory
import java.time.LocalDate

@Entity
class Recruitment(

    @OneToMany(mappedBy = "recruitment", cascade = [CascadeType.ALL], orphanRemoval = true)
    val recruitmentPositions: MutableList<RecruitmentPosition> = mutableListOf(),

    @OneToMany(mappedBy = "recruitment", cascade = [CascadeType.ALL], orphanRemoval = true)
    val recruitmentTechStacks: MutableList<RecruitmentTechStack> = mutableListOf(),

    @Lob
    val content: String,

    val pageViews: Long = 0L,

    val subject: String,

    val contract: String,

    val recruitmentDeadLine: LocalDate,

    val progressPeriod: Int,

    val numberOfPeople: Int,

    @Enumerated(EnumType.STRING)
    val progressMethods: ProgressMethods,

    @Enumerated(EnumType.STRING)
    val recruitmentCategory: RecruitmentCategory,

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long? = null,
): BaseEntity()
