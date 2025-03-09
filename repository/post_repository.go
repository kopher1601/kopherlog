package repository

import (
	"context"
	"kopherlog/db"
	"kopherlog/domain"
	"kopherlog/ent"
	"kopherlog/ent/post"
	"log"
)

type PostRepository interface {
	Save(ctx context.Context, post *domain.Post) (*domain.Post, error)
	FindAll(ctx context.Context, search *domain.PostSearch) ([]*ent.Post, error)
	Delete(ctx context.Context, id int) error
	DeleteAll(ctx context.Context) error
	FindByID(ctx context.Context, id int) (*ent.Post, error)
	SaveAll(ctx context.Context, creates []*domain.PostCreate) error
	Update(ctx context.Context, target *ent.Post, source *domain.Post) error
}

type postRepository struct {
	ent *ent.Client
}

func NewPostRepository(ent *ent.Client) PostRepository {
	return &postRepository{ent: ent}
}

func (p *postRepository) Save(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	var response domain.Post
	err := db.WithTx(ctx, p.ent, func(tx *ent.Tx) error {
		savedPost, err := tx.Post.Create().SetTitle(post.Title).SetContent(post.Content).Save(ctx)
		if err != nil {
			return err
		}
		response = domain.Post{
			ID:      savedPost.ID,
			Title:   savedPost.Title,
			Content: savedPost.Content,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *postRepository) FindAll(ctx context.Context, search *domain.PostSearch) ([]*ent.Post, error) {
	if search == nil {
		search = &domain.PostSearch{
			Page: 0,
			Size: 10,
		}
	}
	posts, err := p.ent.Post.Query().
		Offset(search.Offset()).
		Limit(search.Limit()).
		Order(ent.Desc(post.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postRepository) FindByID(ctx context.Context, id int) (*ent.Post, error) {
	foundPost, err := p.ent.Post.Query().Where(post.ID(id)).First(ctx)
	if err != nil {
		log.Println("postRepository.FindByID:", err)
		return nil, err
	}
	return foundPost, nil
}

func (p *postRepository) SaveAll(ctx context.Context, posts []*domain.PostCreate) error {
	return db.WithTx(ctx, p.ent, func(tx *ent.Tx) error {
		_, err := tx.Post.MapCreateBulk(posts, func(postCreate *ent.PostCreate, i int) {
			postCreate.SetTitle(posts[i].Title).SetContent(posts[i].Content)
		}).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

func (p *postRepository) Update(ctx context.Context, target *ent.Post, source *domain.Post) error {
	_, err := p.ent.Post.UpdateOne(target).
		SetTitle(source.Title).
		SetContent(source.Content).
		Save(ctx)
	if err != nil {
		log.Println("postRepository.Update:", err)
		return err
	}
	return nil
}

func (p *postRepository) Delete(ctx context.Context, id int) error {
	_, err := p.ent.Post.Delete().Where(post.ID(id)).Exec(ctx)
	if err != nil {
		log.Println("postRepository.Delete:", err)
		return err
	}
	return nil
}

func (p *postRepository) DeleteAll(ctx context.Context) error {
	_, err := p.ent.Post.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
