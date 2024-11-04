package jp.co.kopher1601.topu.recruitment.repository

import jp.co.kopher1601.topu.recruitment.domain.Recruitment
import org.springframework.data.jpa.repository.JpaRepository

interface RecruitmentRepository: JpaRepository<Recruitment, Long> {
}