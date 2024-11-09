package jp.co.kopher1601.topu.recruitment.service

import jp.co.kopher1601.topu.recruitment.domain.*
import jp.co.kopher1601.topu.recruitment.repository.PositionRepository
import jp.co.kopher1601.topu.recruitment.repository.RecruitmentRepository
import jp.co.kopher1601.topu.recruitment.repository.TechStackRepository
import jp.co.kopher1601.topu.recruitment.service.dto.PostRecruitment
import jp.co.kopher1601.topu.recruitment.service.dto.RecruitmentResponse
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.data.crossstore.ChangeSetPersister.NotFoundException
import org.springframework.data.repository.findByIdOrNull
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional

@Service
class RecruitmentServiceImpl @Autowired constructor(
    private val recruitmentRepository: RecruitmentRepository,
    private val techStackRepository: TechStackRepository,
    private val positionRepository: PositionRepository,
) : RecruitmentService {

    @Transactional
    override fun post(postRecruitment: PostRecruitment) {

        val recruitment = Recruitment(
            recruitmentCategory = postRecruitment.recruitmentCategory!!,
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
            recruitmentTechStack.makeRelationship()
        }

        postRecruitment.positions?.map { positionName ->
            positionRepository.findByPositionName(positionName) ?: Position(positionName = positionName)
        }?.forEach { position ->
            val recruitmentPosition = RecruitmentPosition(recruitment, position)
            recruitmentPosition.makeRelationship()
        }

        recruitmentRepository.save(recruitment)
    }

    override fun getRecruitment(recruitmentId: Long): RecruitmentResponse {
        val foundRecruitment = recruitmentRepository.findByIdOrNull(recruitmentId)
            ?: throw NotFoundException()

        return RecruitmentResponse.from(foundRecruitment)
    }
}