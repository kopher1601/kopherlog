package jp.co.kopher1601.topu.recruitment.domain

import jakarta.persistence.*

@Entity
class Recruitment(

    @OneToMany(mappedBy = "recruitment", cascade = [CascadeType.ALL], orphanRemoval = true)
    val recruitmentTechStacks: MutableList<RecruitmentTechStack>? = mutableListOf(),

    @Enumerated(EnumType.STRING)
    private val recruitmentCategories: RecruitmentCategories,

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long? = null,
) {

}
