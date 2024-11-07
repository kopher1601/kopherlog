package jp.co.kopher1601.topu.recruitment.repository

import jp.co.kopher1601.topu.recruitment.domain.Position
import org.springframework.data.jpa.repository.JpaRepository

interface PositionRepository: JpaRepository<Position, Long> {

    /**
     * select p
     * from Position p
     * where p.positionName = :positionName
     */
    fun findByPositionName(positionName: String): Position?
}