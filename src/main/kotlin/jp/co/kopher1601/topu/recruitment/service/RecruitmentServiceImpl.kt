package jp.co.kopher1601.topu.recruitment.service

import jp.co.kopher1601.topu.recruitment.repository.RecruitmentRepository
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class RecruitmentServiceImpl @Autowired constructor(
    private val recruitmentRepository: RecruitmentRepository,
) : RecruitmentService {
    override fun post(recruitment: PostRecruitment) {
    }
}