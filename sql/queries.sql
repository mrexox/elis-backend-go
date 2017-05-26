-- Home / Blog
SELECT p.id, p.name, p.content, p.permalink, p.created_at, likes.likes
FROM post p LEFT OUTER JOIN (SELECT post_id, COUNT(ip) AS  likes
                      FROM liker
                      GROUP BY post_id) likes ON p.id = likes.post_id
ORDER BY created_at
LIMIT 4;

-- Blog/:permalink
SELECT p.id, p.name, p.content, p.permalink, p.created_at, likes.likes
FROM post p LEFT OUTER JOIN (SELECT post_id, COUNT(ip) AS  likes
                      FROM liker
                      GROUP BY post_id) likes ON p.id = likes.post_id
WHERE permalink = ?;
-- Tags in blog
SELECT t.name 
FROM post_tag pt INNER JOIN tag t ON pt.tag_id = t.id
WHERE pt.post_id = ?;

