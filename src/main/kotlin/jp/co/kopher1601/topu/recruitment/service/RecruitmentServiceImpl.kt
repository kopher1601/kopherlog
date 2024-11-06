package jp.co.kopher1601.topu.recruitment.service

import jp.co.kopher1601.topu.recruitment.domain.Recruitment
import jp.co.kopher1601.topu.recruitment.domain.RecruitmentTechStack
import jp.co.kopher1601.topu.recruitment.domain.TechStack
import jp.co.kopher1601.topu.recruitment.repository.RecruitmentRepository
import jp.co.kopher1601.topu.recruitment.repository.TechStackRepository
import jp.co.kopher1601.topu.recruitment.service.dto.PostRecruitment
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional

@Service
class RecruitmentServiceImpl @Autowired constructor(
    private val recruitmentRepository: RecruitmentRepository,
    private val techStackRepository: TechStackRepository,
) : RecruitmentService {

    @Transactional
    override fun post(postRecruitment: PostRecruitment) {

        val recruitment = Recruitment(
            recruitmentCategories = postRecruitment.recruitmentCategories!!,
            progressMethods = postRecruitment.progressMethods!!,
            content = postRecruitment.content!!,
            subject = postRecruitment.subject!!,
            contract = postRecruitment.contract!!,
            recruitmentDeadLine = postRecruitment.recruitmentDeadline!!,
            progressPeriod = postRecruitment.progressPeriod!!,
            numberOfPeople = postRecruitment.numberOfPeople!!,
        )

        postRecruitment.techStacks?.map { techStack ->
            techStackRepository.findByTechnologyName(techStack)
                ?: TechStack(technologyName = techStack)
        }?.forEach { techStack ->
            val recruitmentTechStack = RecruitmentTechStack(recruitment, techStack)
            recruitmentTechStack.makeRelationShip()
        }

        recruitmentRepository.save(recruitment)
    }
}