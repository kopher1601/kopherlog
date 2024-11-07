package jp.co.kopher1601.topu.recruitment.domain

import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.GeneratedValue
import jakarta.persistence.GenerationType
import jakarta.persistence.Id
import jakarta.persistence.OneToMany

@Entity
class TechStack(

    @OneToMany(mappedBy = "techStack")
    val recruitmentTechStacks: MutableList<RecruitmentTechStack> = mutableListOf(),

    @Column(unique = true, nullable = false, length = 50)
    private val technologyName: String,

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long? = null,
): BaseEntity()