package jp.co.kopher1601.topu.recruitment.domain

import jakarta.persistence.*

@Entity
class Recruitment(

    @Enumerated(EnumType.STRING)
    private val recruitmentCategories: RecruitmentCategories,

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long? = null,
) {

}
