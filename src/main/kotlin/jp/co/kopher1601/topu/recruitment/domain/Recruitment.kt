package jp.co.kopher1601.topu.recruitment.domain

import jakarta.persistence.*
import java.time.LocalDate

@Entity
class Recruitment(

    @OneToMany(mappedBy = "recruitment", cascade = [CascadeType.ALL], orphanRemoval = true)
    val recruitmentTechStacks: MutableList<RecruitmentTechStack>? = mutableListOf(),

    @Lob
    private val content: String,

    private val pageViews: Long = 0L,

    private val subject: String,

    private val contract: String,

    private val recruitmentDeadLine: LocalDate,

    private val progressPeriod: Int,

    private val numberOfPeople: Int,

    @Enumerated(EnumType.STRING)
    private val progressMethods: ProgressMethods,

    @Enumerated(EnumType.STRING)
    private val recruitmentCategories: RecruitmentCategories,

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long? = null,
) {

}
