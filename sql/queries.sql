-- Posts
SELECT p.id, p.name, p.content, p.permalink, p.visible, p.created_at, IFNULL(likes.likes, 0), p.cover
FROM post p LEFT OUTER JOIN (SELECT post_id, COUNT(ip) AS  likes
                      FROM liker
                      GROUP BY post_id) likes ON p.id = likes.post_id
ORDER BY created_at;

-- Blog/:permalink
SELECT p.id, p.name, p.content, p.permalink, p.created_at, IFNULL(likes.likes, 0), p.cover
FROM post p LEFT OUTER JOIN (SELECT post_id, COUNT(ip) AS  likes
                      FROM liker
                      GROUP BY post_id) likes ON p.id = likes.post_id
WHERE permalink = ? AND visible = TRUE;

-- Tags in blog
SELECT t.name 
FROM post_tag pt INNER JOIN tag t ON pt.tag_id = t.id
WHERE pt.post_id = ?;

-- Portfolio images
SELECT im.id, im.url
FROM portfolio_image pi INNER JOIN image im ON pi.id = im.id;
-- Get cover for giver ID
SELECT id, url
FROM image
WHERE id = ?
