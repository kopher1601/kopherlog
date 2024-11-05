package jp.co.kopher1601.topu.recruitment.service

import jp.co.kopher1601.topu.recruitment.domain.Recruitment
import jp.co.kopher1601.topu.recruitment.domain.RecruitmentTechStack
import jp.co.kopher1601.topu.recruitment.domain.TechStack
import jp.co.kopher1601.topu.recruitment.repository.RecruitmentRepository
import jp.co.kopher1601.topu.recruitment.repository.TechStackRepository
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class RecruitmentServiceImpl @Autowired constructor(
    private val recruitmentRepository: RecruitmentRepository,
    private val techStackRepository: TechStackRepository,
) : RecruitmentService {
    override fun post(postRecruitment: PostRecruitment) {

        val recruitment = Recruitment(recruitmentCategories = postRecruitment.recruitmentCategories!!)

        postRecruitment.techStacks?.map { techStack ->
            techStackRepository.findByTechnologyName(techStack)
                ?: TechStack(technologyName = techStack)
        }?.forEach { techStack ->
            val recruitmentTechStack = RecruitmentTechStack(recruitment, techStack)
            recruitmentTechStack.makeRelationShip()
        }

    }
}