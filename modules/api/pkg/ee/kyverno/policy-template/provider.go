//go:build ee

/*
                  Kubermatic Enterprise Read-Only License
                         Version 1.0 ("KERO-1.0”)
                     Copyright © 2025 Kubermatic GmbH

   1.	You may only view, read and display for studying purposes the source
      code of the software licensed under this license, and, to the extent
      explicitly provided under this license, the binary code.
   2.	Any use of the software which exceeds the foregoing right, including,
      without limitation, its execution, compilation, copying, modification
      and distribution, is expressly prohibited.
   3.	THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
      EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
      MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
      IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
      CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
      TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
      SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

   END OF TERMS AND CONDITIONS
*/

package policytemplate

import (
	"context"
	"fmt"

	"k8c.io/dashboard/v2/pkg/provider"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type PolicyTemplateProvider struct {
	privilegedClient ctrlruntimeclient.Client
}

func NewPolicyTemplateProvider(privilegedClient ctrlruntimeclient.Client) *PolicyTemplateProvider {
	return &PolicyTemplateProvider{
		privilegedClient: privilegedClient,
	}
}

func (p *PolicyTemplateProvider) CreateUnsecured(ctx context.Context, policyTemplate *kubermaticv1.PolicyTemplate) (*kubermaticv1.PolicyTemplate, error) {
	if err := p.privilegedClient.Create(ctx, policyTemplate); err != nil {
		return nil, err
	}

	return policyTemplate, nil
}

func (p *PolicyTemplateProvider) ListUnsecured(ctx context.Context) (*kubermaticv1.PolicyTemplateList, error) {
	policyTemplateList := &kubermaticv1.PolicyTemplateList{}
	if err := p.privilegedClient.List(ctx, policyTemplateList); err != nil {
		return nil, err
	}

	return policyTemplateList, nil
}

func (p *PolicyTemplateProvider) GetUnsecured(ctx context.Context, policyTemplateName string) (*kubermaticv1.PolicyTemplate, error) {
	policyTemplate := &kubermaticv1.PolicyTemplate{}
	if err := p.privilegedClient.Get(ctx, ctrlruntimeclient.ObjectKey{Name: policyTemplateName}, policyTemplate); err != nil {
		return nil, err
	}

	return policyTemplate, nil
}

func (p *PolicyTemplateProvider) PatchUnsecured(ctx context.Context, user *provider.UserInfo, updatedpolicyTemplate *kubermaticv1.PolicyTemplate) (*kubermaticv1.PolicyTemplate, error) {
	client := p.privilegedClient

	existing := &kubermaticv1.PolicyTemplate{}

	if err := client.Get(ctx, ctrlruntimeclient.ObjectKey{Name: updatedpolicyTemplate.Name}, existing); err != nil {
		return nil, err
	}

	updated := existing.DeepCopy()
	updated.Spec = updatedpolicyTemplate.Spec
	if updated.Spec.ProjectID == existing.Spec.ProjectID || user.IsAdmin {
		if err := client.Patch(ctx, updated, ctrlruntimeclient.MergeFrom(existing)); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user %s is not allowed to update the policy template %s", user.Email, updatedpolicyTemplate.Name)
	}

	return updated, nil
}

func (p *PolicyTemplateProvider) DeleteUnsecured(ctx context.Context, policyTemplateName string, projectID string, user *provider.UserInfo) error {
	client := p.privilegedClient

	existing := &kubermaticv1.PolicyTemplate{}
	if err := client.Get(ctx, ctrlruntimeclient.ObjectKey{Name: policyTemplateName}, existing); err != nil {
		return err
	}
	if existing.Spec.ProjectID == projectID || (user.IsAdmin && projectID == "") {
		if err := client.Delete(ctx, existing); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("user %s is not allowed to delete the policy template %s", user.Email, policyTemplateName)
	}
	return nil
}
