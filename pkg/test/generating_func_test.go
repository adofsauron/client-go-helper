package test

import (
	"client-go-helper/pkg/kubectl"
	"fmt"
	"testing"
)

func TestWriteDeleteCode(t *testing.T)  {

	clientset, _ := kubectl.InitClient()

	//输出转输入
	writeDeleteCode(GetFunctionName2(clientset.BatchV1beta1()),"BatchV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.AuditregistrationV1alpha1()),"AuditregistrationV1alpha1")

	writeDeleteCode(GetFunctionName2(clientset.AdmissionregistrationV1beta1()),"AdmissionregistrationV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.BatchV1()),"BatchV1")

	writeDeleteCode(GetFunctionName2(clientset.CertificatesV1beta1()),"CertificatesV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.RbacV1()),"RbacV1")

	writeDeleteCode(GetFunctionName2(clientset.CoordinationV1beta1()),"CoordinationV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.NetworkingV1()),"NetworkingV1")

	writeDeleteCode(GetFunctionName2(clientset.AppsV1()),"AppsV1")

	writeDeleteCode(GetFunctionName2(clientset.AutoscalingV2beta2()),"AutoscalingV2beta2")

	writeDeleteCode(GetFunctionName2(clientset.PolicyV1beta1()),"PolicyV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.AuthorizationV1()),"AuthorizationV1")

	writeDeleteCode(GetFunctionName2(clientset.CoreV1()),"CoreV1")

	writeDeleteCode(GetFunctionName2(clientset.SettingsV1alpha1()),"SettingsV1alpha1")

	writeDeleteCode(GetFunctionName2(clientset.AdmissionregistrationV1alpha1()),"AdmissionregistrationV1alpha1")

	writeDeleteCode(GetFunctionName2(clientset.AuthenticationV1()),"AuthenticationV1")

	writeDeleteCode(GetFunctionName2(clientset.ExtensionsV1beta1()),"ExtensionsV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.SchedulingV1beta1()),"SchedulingV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.StorageV1()),"StorageV1")
}



func writeDeleteCode (args map[string]bool,api string){
	for  v,vbool := range args {
		if v!= "" {
			var vs = v + "s"
			if  v[len(v)-1:]== "s"  {
				vs = v + "es"
			}
			if v[len(v)-1:]== "y" {
				vs = v[:len(v)-1] + "ies"
			}

			var ns = "i.Namespace"
			if vbool == false {
				ns = ""
			}
			fmt.Printf(`
func (i *%s) Delete (opts *v1.DeleteOptions) (error) {
	e := clientset.%s().%s(%s).Delete(i.Name,opts)
	if e != nil {
		return e
	}
	return nil
}
`, v, api, vs, ns)
		}
	}


}