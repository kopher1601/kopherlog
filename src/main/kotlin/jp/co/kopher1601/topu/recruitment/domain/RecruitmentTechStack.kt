package jp.co.kopher1601.topu.recruitment.domain

import jakarta.persistence.*

@Entity
class RecruitmentTechStack(

    @ManyToOne(fetch = FetchType.LAZY, cascade = [CascadeType.PERSIST])
    private val recruitment: Recruitment,

    @ManyToOne(fetch = FetchType.LAZY, cascade = [CascadeType.PERSIST])
    val techStack: TechStack,

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long? = null,
): BaseEntity() {

    fun makeRelationship() {
        this.recruitment.recruitmentTechStacks?.add(this)
        this.techStack.recruitmentTechStacks.add(this)
    }
}