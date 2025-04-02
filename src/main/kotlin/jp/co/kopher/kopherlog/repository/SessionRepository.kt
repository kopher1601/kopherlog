package jp.co.kopher.kopherlog.repository

import jp.co.kopher.kopherlog.domain.Session
import org.springframework.data.repository.CrudRepository

interface SessionRepository : CrudRepository<Session, Long>