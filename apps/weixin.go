package apps

import (
	"github.com/Michael2008S/social-auth"
)

type Weixin struct {
	BaseProvider
}

func (p *Weixin) GetType() social.SocialType {
	return social.SocialWeixin
}

func (p *Weixin) GetName() string {
	return "Weixin"
}

func (p *Weixin) GetPath() string {
	return "weixin"
}

func (p *Weixin) GetIndentify(tok *social.Token) (string, error) {
	return tok.GetExtra("openid"), nil
}

var _ social.Provider = new(Weixin)

func NewWeixin(clientId, secret string) *Weixin {
	p := new(Weixin)
	p.App = p
	p.ClientId = clientId
	p.ClientSecret = secret
	p.Scope = "snsapi_login"
	p.AuthURL = "https://open.weixin.qq.com/connect/qrconnect"
	p.TokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token"
	p.RedirectURL = social.DefaultAppUrl + "login/weixin/access"
	p.AccessType = "offline"
	p.ApprovalPrompt = "auto"
	return p
}
