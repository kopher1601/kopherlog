package jp.co.kopher1601.topu.recruitment.domain

import jakarta.persistence.*

@Entity
class RecruitmentPosition(

    @ManyToOne(fetch = FetchType.LAZY, cascade = [CascadeType.PERSIST])
    private val recruitment: Recruitment,

    @ManyToOne(fetch = FetchType.LAZY, cascade = [CascadeType.PERSIST])
    private val position: Position,

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private val id: Long? = null,
): BaseEntity() {
    fun makeRelationship() {
        this.position.recruitmentPositions.add(this)
        this.recruitment.recruitmentPositions.add(this)
    }
}