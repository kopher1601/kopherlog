package jp.co.kopher1601.topu.recruitment.repository

import jp.co.kopher1601.topu.recruitment.domain.TechStack
import org.springframework.data.jpa.repository.JpaRepository

interface TechStackRepository: JpaRepository<TechStack, Long> {

    /**
     * select t
     * from TechStack t
     * where t.technologyName := techStack
     */
    fun findByTechnologyName(techStack: String): TechStack?
}