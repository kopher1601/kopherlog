package jp.co.kopher1601.topu.recruitment.domain

import jakarta.persistence.*

@Entity
class RecruitmentTechStack(

    @ManyToOne(fetch = FetchType.LAZY)
    private val recruitment: Recruitment,

    @ManyToOne(fetch = FetchType.LAZY)
    private val techStack: TechStack,

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long? = null,
) {

    fun makeRelationShip() {
        this.recruitment.recruitmentTechStacks?.add(this)
        this.techStack.recruitmentTechStacks.add(this)
    }
}