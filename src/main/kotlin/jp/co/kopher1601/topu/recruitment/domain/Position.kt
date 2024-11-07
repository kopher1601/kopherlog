package jp.co.kopher1601.topu.recruitment.domain

import jakarta.persistence.*

@Entity
class Position(

    @OneToMany(mappedBy = "position")
    val recruitmentPositions: MutableList<RecruitmentPosition> = mutableListOf(),

    @Column(nullable = false, unique = true, length = 50)
    private val positionName: String,

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long? = null,
): BaseEntity()