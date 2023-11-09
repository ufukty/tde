package runner_communicator

import (
	"go/ast"
	"tde/internal/evolution/models"

	"fmt"
	"testing"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

func Test_BatchDivide(t *testing.T) {
	checkBatchSize := func(batch *Batch, subBatch *Batch, noBatches int) error {
		maxExpected := len(batch.Subjects)/noBatches + 1
		minExpected := len(batch.Subjects) / noBatches

		if len(subBatch.Subjects) > maxExpected || len(subBatch.Subjects) < minExpected {
			return fmt.Errorf("batch size out of range, expected %d <= size <= %d, got %d", minExpected, maxExpected, len(subBatch.Subjects))
		}
		return nil
	}

	checkAllItemsPlaced := func(batch *Batch, subBatches []*Batch) error {
		merged := Batch{}
		for _, subBatch := range subBatches {
			merged.Subjects = append(merged.Subjects, subBatch.Subjects...)
		}

		for _, item := range batch.Subjects {
			if slices.Index(merged.Subjects, item) == -1 {
				return fmt.Errorf("None of the subbatches has the item: %v", item)
			}
		}

		return nil
	}

	checkItems := func(batch *Batch, subBatch *Batch) error {
		for _, c := range subBatch.Subjects {
			found := false
			for _, cc := range batch.Subjects {
				if c == cc {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("subject %v not found in original batch", c)
			}
		}
		return nil
	}

	batch := &Batch{
		File: &ast.File{
			Name: &ast.Ident{Name: "blabla"},
		},
		Subjects: []*models.Subject{
			{Sid: models.Sid("0a40f3cd-5d02-594c-bf1a-6c5aca05edb8")},
			{Sid: models.Sid("2901cbd4-72e7-5b97-9fa0-d50f863f2d0b")},
			{Sid: models.Sid("c33efc69-0130-5301-b242-1dbe83e53e42")},
			{Sid: models.Sid("31bc5f21-67f8-5088-bc52-4056295fe96a")},
			{Sid: models.Sid("5f074b1d-3b91-50cd-a11e-e8a5e8f1ac36")},
			{Sid: models.Sid("32817d78-2e71-57b8-8c3c-5ff704b17b42")},
			{Sid: models.Sid("c0439809-9087-50ed-a296-a490e615cde0")},
			{Sid: models.Sid("1efdb10e-8693-5992-9435-fd958bfa372d")},
			{Sid: models.Sid("9525d2fb-c448-5db3-9983-ecb9ac5345d7")},
			{Sid: models.Sid("3b1108f2-c4fe-5cde-a2d7-0360caebdd5a")},
			{Sid: models.Sid("b1e67897-b964-5440-8440-ce0af1185a5c")},
			{Sid: models.Sid("c3a26537-0159-525f-868c-9cb7b370b276")},
			{Sid: models.Sid("9c5b6e93-ab69-5ff2-bbde-d9f49e2dcd63")},
			{Sid: models.Sid("4a5024ec-a546-5c5b-9835-d191754df6b1")},
			{Sid: models.Sid("60ea077b-6789-531d-a967-c226bb7fa69f")},
			{Sid: models.Sid("e971e216-873b-5907-8d6b-76171c886405")},
			{Sid: models.Sid("f2da10ed-a825-55b2-ac52-0670aae41f56")},
			{Sid: models.Sid("805f6fe1-486d-5386-90bd-4ec492e0e2b9")},
			{Sid: models.Sid("d2481f8a-264a-5d26-832e-adf73dac328b")},
			{Sid: models.Sid("144e1c30-1ee3-577e-9c5a-bc1435ab6b06")},
			{Sid: models.Sid("e1b7e032-e13b-5991-916d-eec78fe02543")},
			{Sid: models.Sid("5aa14327-8395-509f-a500-18634f134bc7")},
			{Sid: models.Sid("adc3fa48-afa6-5171-9043-bb303664a2f3")},
			{Sid: models.Sid("8912df97-9120-5951-b2da-d3749ccdfccf")},
			{Sid: models.Sid("a5705dec-3529-5381-8d42-38fbcd364f61")},
			{Sid: models.Sid("eccd8712-618e-5144-8daf-732168ef9c31")},
			{Sid: models.Sid("7afd0398-d554-502f-ba6d-5fc0c415e5ea")},
			{Sid: models.Sid("9ae3d665-18bf-5206-97a5-ce6316e675a9")},
			{Sid: models.Sid("be60b933-1fe6-53cf-9c95-8237e87ac899")},
			{Sid: models.Sid("fec310c4-0fc0-51ba-b2bf-4990ca3b8695")},
			{Sid: models.Sid("baf80a78-f07e-5308-af59-27b62003d6ee")},
			{Sid: models.Sid("be051481-c079-598a-b28a-9fd2bd899494")},
			{Sid: models.Sid("7cbaade0-39c9-5298-bc11-521e3b947fc0")},
			{Sid: models.Sid("5fd9567e-04de-50ef-ba3f-a2194869b0ab")},
			{Sid: models.Sid("dc9b920d-7734-57a3-9183-0599f2aaf2b1")},
			{Sid: models.Sid("9a23f972-a3b4-56c5-87f6-e23d75ab8085")},
			{Sid: models.Sid("b534d245-4c24-57fb-ae51-2520d60e675d")},
			{Sid: models.Sid("56188b2e-834e-557f-819e-345ece3eaf7b")},
			{Sid: models.Sid("15317383-eb25-5001-99d8-0c93185c3138")},
			{Sid: models.Sid("a28b1f49-84de-5a15-b756-a1bf521e116b")},
			{Sid: models.Sid("4adf206c-d58d-5185-8808-1f9785985b98")},
			{Sid: models.Sid("ee24fb34-846e-57b6-9015-e915317b350e")},
			{Sid: models.Sid("f58272e2-37a1-5614-b52d-0c3b6ca0bc4c")},
			{Sid: models.Sid("a8ea22fd-ea3c-5e9e-8a6a-ea3df044467a")},
			{Sid: models.Sid("36e5accd-5d41-53b6-8df2-1cd8ef10dfc1")},
			{Sid: models.Sid("e06e593a-14b0-5070-9b2c-e0ceb073f4be")},
			{Sid: models.Sid("69e79178-0b02-5a2b-8a9f-10e4d1e5672e")},
			{Sid: models.Sid("13f27439-a1ee-5c40-879d-b07fa1466a38")},
			{Sid: models.Sid("b312f0e0-af9f-535c-91a2-89b9ae1bf5fa")},
			{Sid: models.Sid("883478ee-c1ef-51cd-a576-3daaf003fdb7")},
			{Sid: models.Sid("4b4c57dd-82d7-5054-a92a-d3ed643c0196")},
			{Sid: models.Sid("41b737e1-923a-5802-8544-449c7bcea851")},
			{Sid: models.Sid("c4dd644e-8fc7-5fd8-bd45-4165416896bc")},
			{Sid: models.Sid("b107c4e8-feba-5745-a5c9-14166b313915")},
			{Sid: models.Sid("56a50afc-57f1-5837-b2da-3563b5c59fbe")},
			{Sid: models.Sid("e75a719b-3966-52c0-82a4-0e728a353d1c")},
			{Sid: models.Sid("eef87cf7-acde-5fd1-b052-209f8e8e7348")},
			{Sid: models.Sid("d5881b1d-cd93-55b0-a788-eb7285883e45")},
			{Sid: models.Sid("8f6dc741-5a65-5d83-b0f5-feb21f72c432")},
			{Sid: models.Sid("f2ab924e-2fc5-55f9-b748-add728861457")},
			{Sid: models.Sid("b0098db4-b2cd-5922-b6cc-82b8137ae3e8")},
			{Sid: models.Sid("2d1bd358-540c-5ee6-81fb-f0bc4837e3ea")},
			{Sid: models.Sid("b66cc41a-40bc-5cc7-8fd7-328028906925")},
			{Sid: models.Sid("f68f5375-3205-5eb4-aa65-862f8290de69")},
			{Sid: models.Sid("d0ddd3c9-56c1-5db0-9196-916a92d18e62")},
			{Sid: models.Sid("5547d16c-fb50-519e-85ce-ce66ae815360")},
			{Sid: models.Sid("3beb9680-5cad-550c-9cb5-e98a4109a560")},
			{Sid: models.Sid("a4d2aea9-a629-5f7b-a721-6d317f95a59d")},
			{Sid: models.Sid("e9a53135-1a89-5d84-b717-7eff6343339b")},
			{Sid: models.Sid("9f19e3fa-22b2-5705-827d-0c2cd6ae37bb")},
			{Sid: models.Sid("10b6a83d-c9bb-57b1-92e7-ac40efd0a87f")},
			{Sid: models.Sid("e1b83767-c208-5fb5-8cf4-749f1070f9c7")},
			{Sid: models.Sid("7e4f9969-1462-59d8-8fac-035c47fab00e")},
			{Sid: models.Sid("4d155049-0601-5c01-a846-6e16391ada6b")},
			{Sid: models.Sid("28c2499a-306c-5945-9ed8-394fc98f019f")},
			{Sid: models.Sid("46a06cfa-db5f-5944-839a-91b34841c9e7")},
			{Sid: models.Sid("4ad6ee04-f28f-5872-82d5-8284bf5cc8f2")},
			{Sid: models.Sid("a2fe346f-5c84-51dd-a67d-16cc3fc2c6c7")},
			{Sid: models.Sid("b24a8a0a-262a-546f-97c5-5656c1fbb8f8")},
			{Sid: models.Sid("575e4ab7-597f-500b-b9d4-c5ffc2735047")},
			{Sid: models.Sid("ee87a6a5-a050-5483-964e-205caadd131d")},
			{Sid: models.Sid("bc59aeb5-c093-5a2b-8954-9917eda3e51d")},
			{Sid: models.Sid("c5ee728e-e040-571d-92c5-5347b87ecb42")},
			{Sid: models.Sid("c5938cd7-5a21-5f84-8e96-c2115f27645b")},
			{Sid: models.Sid("4bca3803-2e2a-5443-aaf8-3bae4db0fbf6")},
			{Sid: models.Sid("77522330-93c2-5f52-8f8b-ec224f9f7289")},
			{Sid: models.Sid("e66fc896-207e-552b-8f73-d49b3c19e367")},
			{Sid: models.Sid("c4954209-89a2-571b-ac94-5705cc7e1883")},
			{Sid: models.Sid("286bd405-792a-5c82-9627-e27c102113d0")},
			{Sid: models.Sid("6faea496-ba7a-5080-a284-a93251b43b7a")},
			{Sid: models.Sid("d74d5dfd-a7f4-5dd3-a267-d9dc7c59af2d")},
			{Sid: models.Sid("b55e3353-13d6-5070-a05f-6fea4334608b")},
			{Sid: models.Sid("d2f284e6-9947-5e8e-bc8d-4a1b1354afb3")},
			{Sid: models.Sid("f594d589-49b7-5f52-b249-3661bbfdd049")},
			{Sid: models.Sid("17cc9928-766f-582b-a4df-7f1dd8d192a1")},
			{Sid: models.Sid("5c6fe1b4-33d7-51d1-aae8-c5877518e776")},
			{Sid: models.Sid("0cf27cd5-f742-51f6-ba2c-63841df9361f")},
			{Sid: models.Sid("974b2b04-4631-5c73-acc4-3a7eba66ca99")},
			{Sid: models.Sid("d105d930-eb06-5a16-9620-0971ac10eb2e")},
			{Sid: models.Sid("15168a45-a46d-5ea1-994d-6207ee901adf")},
			{Sid: models.Sid("21325c48-867c-5d9e-bb22-5093c42981db")},
			{Sid: models.Sid("aab4697c-2995-5ec0-b2eb-bbf487a74c5f")},
			{Sid: models.Sid("b48dc5c0-c29e-5e93-b9a6-e26c1e11c7d3")},
			{Sid: models.Sid("8d5300ca-4e1e-59ce-9ee2-67c60965daa3")},
			{Sid: models.Sid("3d689854-5d1a-5379-afd0-45fbaf1f10a1")},
			{Sid: models.Sid("c002a522-0212-523b-bc1a-1549bdda069f")},
			{Sid: models.Sid("6dabf5ae-8c85-57e6-8d28-94acd0ddf403")},
			{Sid: models.Sid("87e9f329-6e54-5db0-bf44-2a609b1dfcbb")},
			{Sid: models.Sid("3a5e0aef-ed28-5048-8ce5-5557e4c3938b")},
			{Sid: models.Sid("50e06961-cd3d-518e-8a69-77da74e91f36")},
			{Sid: models.Sid("04a8cb6d-4ce7-5263-911c-8cc95ed4c5d0")},
			{Sid: models.Sid("bb15b989-eb81-52a7-8478-b885201465fb")},
			{Sid: models.Sid("022625af-331d-5fdf-87bb-5bafbc22bbe1")},
			{Sid: models.Sid("f04977e8-4a66-518d-bed7-9ece284a6c6f")},
			{Sid: models.Sid("350f6111-aa3d-5800-a4a9-7f00f5114e34")},
			{Sid: models.Sid("fb3b0f36-368c-5bc1-85b5-141414c9924b")},
			{Sid: models.Sid("6e693c79-93c9-5aa5-b8da-5ec75dbabf95")},
			{Sid: models.Sid("2e6c96e3-0b62-5d89-ac33-4517c004361c")},
			{Sid: models.Sid("c8f92e45-3a96-5376-8a7d-aa9de564a72d")},
			{Sid: models.Sid("a64915c9-8d96-500d-9bdf-d0649b0da33a")},
			{Sid: models.Sid("f972d8fe-0c6a-5fe5-a535-e0050ae6072d")},
			{Sid: models.Sid("dcea01a6-ea0c-5ab3-9f2f-324403a5ec77")},
			{Sid: models.Sid("a8dd4574-e90a-595e-a71f-84fbfb561712")},
			{Sid: models.Sid("b83451bb-494a-5a4c-b888-4d4175d5c37a")},
			{Sid: models.Sid("72c8d346-186a-52e6-976a-4885fdc09cc8")},
			{Sid: models.Sid("90dcb845-9740-5e29-8f13-76e48884cb0a")},
			{Sid: models.Sid("a2ffee5a-4948-5104-9a44-d76d12010a66")},
			{Sid: models.Sid("679df4cb-dba8-58df-a234-ba7ab5f68512")},
			{Sid: models.Sid("2f968ef3-63a0-5865-a8a4-144fe97de00f")},
			{Sid: models.Sid("c0bb3325-d38e-5126-822e-bb4339645e08")},
			{Sid: models.Sid("c273fd73-d3c7-5f0d-a4b8-5b2b9a94d133")},
			{Sid: models.Sid("73ad205f-5f20-5cf6-b0b8-4af69998ef8e")},
			{Sid: models.Sid("f68f3efd-aac8-5c6f-94f0-f32549dba2bd")},
			{Sid: models.Sid("0a46e858-5500-5342-b5a1-251d8c8bec47")},
			{Sid: models.Sid("f42058fc-513e-56db-821d-738ac8695e81")},
			{Sid: models.Sid("d8992870-9cb7-55e0-9794-2aa857c08efb")},
			{Sid: models.Sid("d362b622-3ba8-5a35-bcdd-6aed1eff0f8b")},
			{Sid: models.Sid("99d322ae-cf03-5011-8324-fb7f7faa4eb6")},
			{Sid: models.Sid("6b2843d6-7631-5928-a656-8f9aae6c27dc")},
			{Sid: models.Sid("98158575-0013-5b4e-a41a-b216d97ac71c")},
			{Sid: models.Sid("ae2d6ca1-71ae-562f-998d-0194b665d1ab")},
			{Sid: models.Sid("a0112f2b-fb7c-52d4-bd39-4a0ec09fedfc")},
			{Sid: models.Sid("a48df952-c834-5969-b093-dfd7dc51f033")},
			{Sid: models.Sid("a2519c4d-fe3d-5ec3-ade3-5e55265c0a51")},
			{Sid: models.Sid("74ac5484-ff48-5d7b-b1c1-b5f681c6687c")},
			{Sid: models.Sid("a02ca99d-bb45-582d-a132-0224e02e4b73")},
			{Sid: models.Sid("27e0152c-4990-5b1a-ac1a-fa0bb3143b4d")},
			{Sid: models.Sid("1719c34b-de61-59fb-8205-aa36e45a6f1c")},
			{Sid: models.Sid("0a987463-eafe-5302-ba70-b160c0b96b96")},
			{Sid: models.Sid("7deae393-bc39-524d-a331-98ada76b19cb")},
			{Sid: models.Sid("3ece7013-10e3-5f45-8a0a-25e71da6c2da")},
			{Sid: models.Sid("f165878f-ca32-53d8-8cbb-aa43af8d8c3c")},
			{Sid: models.Sid("d6235aa3-d699-50dd-9f2f-2159c99679cd")},
			{Sid: models.Sid("ae98b610-c211-5533-820d-4c764a5aa16f")},
			{Sid: models.Sid("efd562cf-2502-5e2c-8888-0b1cd6519736")},
			{Sid: models.Sid("f69eb48d-e864-5077-acde-56d25300e796")},
			{Sid: models.Sid("9f02b12b-6300-57a1-b3c8-7c0d5271f21b")},
			{Sid: models.Sid("aad64c35-57b9-5e92-985f-aee305a2e4cf")},
			{Sid: models.Sid("ee2f039e-c227-5672-86a3-848eadc24989")},
			{Sid: models.Sid("f4b621f8-ea54-51be-9abe-6cb58b09d3e1")},
			{Sid: models.Sid("4838efba-81cc-5728-b58c-af08f24f1659")},
			{Sid: models.Sid("95df06ca-c4ea-5402-9f29-1464751f8acd")},
			{Sid: models.Sid("3570279c-5f32-546e-8049-5b2f5f61d154")},
			{Sid: models.Sid("512abd39-c8f1-5fce-a1c0-242290da5db6")},
			{Sid: models.Sid("064ae3c0-2bde-5d01-adc5-8f84d0870b3f")},
			{Sid: models.Sid("3ef087c5-8ee7-5004-8467-1e156c736cfc")},
			{Sid: models.Sid("c7d0a1bf-08c0-56df-94ee-a946de92cc89")},
			{Sid: models.Sid("640444ae-941a-5f9b-b1bf-080c209a872f")},
			{Sid: models.Sid("9a02dc00-1b8d-572f-9eba-fc1ae299fb3a")},
			{Sid: models.Sid("88029ccc-9847-519f-ba80-2d95279cc6ca")},
			{Sid: models.Sid("0249864e-e700-520a-89d7-dc14b5d0dc9c")},
			{Sid: models.Sid("f7eb1d3a-d323-588e-864f-08b558f387b9")},
			{Sid: models.Sid("c0dab5bb-08ed-5027-887a-78deabdbec4f")},
			{Sid: models.Sid("a5520932-c935-57ed-93e5-7c4407f4f7b7")},
			{Sid: models.Sid("0627cd46-fb12-5105-9272-33a86c3ff617")},
			{Sid: models.Sid("00c07fac-2c54-5d8c-ab95-006fd57f3db7")},
			{Sid: models.Sid("4c441002-1cf7-58e6-8173-9d77e55644df")},
			{Sid: models.Sid("a3276500-12df-58e3-af27-2be4defcd700")},
			{Sid: models.Sid("7593be2f-a788-590d-b394-8d0ba37e3bd0")},
			{Sid: models.Sid("e7f56405-991f-5a72-934c-b3bb1009327a")},
			{Sid: models.Sid("5531222d-dd90-58e0-940f-dae664749957")},
			{Sid: models.Sid("49fa3495-900c-50fd-89b5-6dfdbfba2fc6")},
			{Sid: models.Sid("65f6f72a-02a6-50bd-9f5f-e50fae823fc8")},
			{Sid: models.Sid("9dea40cf-9031-5c31-ae26-debac8449e1d")},
			{Sid: models.Sid("dd69c590-8a64-54e6-ba55-e33b159a92f9")},
			{Sid: models.Sid("495eecc5-36db-5066-a95a-771e0654eb37")},
			{Sid: models.Sid("2771ff78-ebad-5bd6-a1e8-df42f417d22e")},
			{Sid: models.Sid("22a9fa04-8c9a-5d20-89d1-8248c42d686e")},
			{Sid: models.Sid("2e5052f6-1ee0-5b45-b533-2303ee30e05d")},
			{Sid: models.Sid("0b1ba8d8-f0be-521b-b845-216aa48eb871")},
			{Sid: models.Sid("81a62c4a-2061-5fb9-83b9-bd75ac77f30f")},
			{Sid: models.Sid("f5102884-1117-5059-99d1-1a1f1461b062")},
			{Sid: models.Sid("bbb12c3e-395e-5233-8720-cfd9161f6ccb")},
			{Sid: models.Sid("6f9e562b-e394-59f0-bac4-26b0953c1653")},
			{Sid: models.Sid("97f03b6f-d842-578c-ad2d-2b65ed6645d4")},
			{Sid: models.Sid("b99ae5c2-d232-53eb-ba66-925ecee5e3ea")},
			{Sid: models.Sid("ba394f36-c462-578d-998f-3bd05df73275")},
			{Sid: models.Sid("e2059873-dbee-5009-97c0-8ed9406ebd07")},
			{Sid: models.Sid("df73b8a6-522a-5cdc-9fce-cf1d9fff761e")},
			{Sid: models.Sid("35301f0b-a0ca-5259-916f-7fbf862cce29")},
			{Sid: models.Sid("affbf134-25f6-5b26-8c05-90b757713a8f")},
			{Sid: models.Sid("6e558925-2fc0-56c3-8e0c-811e08090e57")},
			{Sid: models.Sid("a7edf474-e5dd-5995-a925-be41e2fef200")},
			{Sid: models.Sid("3292f169-e8cd-595d-9fac-91d3078f1f15")},
			{Sid: models.Sid("b7720e1c-cfeb-51cb-9ab4-bd10dfe9f76c")},
			{Sid: models.Sid("f7771b88-dd1a-5701-8313-de13fc57700b")},
			{Sid: models.Sid("9c294e5b-b4d6-51f7-937f-bab6fc0ebb86")},
			{Sid: models.Sid("9add59fc-c9c2-598f-83e7-d81a0738fe29")},
			{Sid: models.Sid("c6e62b39-b9f3-5158-8bb5-e3ffb30c7ae8")},
			{Sid: models.Sid("76820be0-9625-5022-a77d-e7e934277b13")},
			{Sid: models.Sid("fe028060-7b49-51eb-9e63-79e8825a47f8")},
			{Sid: models.Sid("36b05b29-bb04-546e-8e59-4c67f40046c4")},
			{Sid: models.Sid("35e18602-412f-5f8f-8921-91d426d588e1")},
			{Sid: models.Sid("26ce239b-e4e0-5220-a1e3-68aa7b0545fd")},
			{Sid: models.Sid("53233f72-7fa7-5b9b-a569-009f54574898")},
			{Sid: models.Sid("52e91c09-a26c-5604-9ede-5a1aba029106")},
			{Sid: models.Sid("0fb13e54-5837-5425-882e-1eb844e4c7cc")},
			{Sid: models.Sid("a30e20a1-f1f3-55d8-bbc5-7ded4bf3d6c7")},
			{Sid: models.Sid("8fdaf5f7-95f4-53cc-a11a-ea2132bdfa6d")},
			{Sid: models.Sid("a61ff288-02d4-518f-9a41-5bcba7afb957")},
			{Sid: models.Sid("163d3ef4-b087-5c1d-b5d4-d4051357150f")},
			{Sid: models.Sid("7d5ff472-9464-5e65-9f52-9cf28907277e")},
			{Sid: models.Sid("c9b69e3a-601e-5398-93a9-169b453b47ee")},
			{Sid: models.Sid("219b6c8c-73d1-5971-a77c-dd0e0ff39df6")},
			{Sid: models.Sid("3c35ef39-aa72-5374-a291-5099bf23e4ee")},
			{Sid: models.Sid("0d9c7701-cf73-572b-9a7c-ef35d8ab472c")},
			{Sid: models.Sid("d712c83f-77d6-504b-b0cf-c5e0d5656b17")},
			{Sid: models.Sid("c948e183-861c-5d73-bddc-d4308b53fc55")},
			{Sid: models.Sid("0e92952a-381f-53b3-9755-ca1b428b17a0")},
			{Sid: models.Sid("97d6e52a-56d0-5698-b6b1-ffba6e74722e")},
			{Sid: models.Sid("8056df46-6305-504f-a19c-51702020c92c")},
			{Sid: models.Sid("c981a644-5805-576c-ab58-481745848e09")},
			{Sid: models.Sid("d38115cd-9596-5961-beae-19805dc1406d")},
			{Sid: models.Sid("a320373e-eb21-5c10-bc56-0798b835907c")},
			{Sid: models.Sid("fd12420a-ce36-5f24-bb3f-7266620d250b")},
			{Sid: models.Sid("6762145b-dccc-50eb-8ea0-22e2bf9e11d7")},
			{Sid: models.Sid("25dd84c4-1f29-5189-b874-4508e1bd632a")},
			{Sid: models.Sid("6a9fd923-8b2f-5f5e-b888-e67fd0816d72")},
			{Sid: models.Sid("06cdba9c-e180-53e3-b691-e0b5f0155f01")},
			{Sid: models.Sid("98552a95-9745-5b20-8a38-41619eee8600")},
			{Sid: models.Sid("23f9b648-fa5f-5b7e-a0e9-8b1d2257efd8")},
			{Sid: models.Sid("cca4886e-c5db-5db8-a56f-6cd9cd311447")},
			{Sid: models.Sid("50cbabdd-0093-54e4-92bc-441402fe5cc8")},
			{Sid: models.Sid("589c922a-aa71-57da-b59e-d1e484f09a3d")},
			{Sid: models.Sid("6f13f7bf-d1bf-5e0e-9c13-949bf0f93958")},
			{Sid: models.Sid("20d75553-321e-5386-bbc6-12267f761b22")},
			{Sid: models.Sid("365c9b20-39c1-5c5e-a39c-7381ba65b7b2")},
			{Sid: models.Sid("eb7f96da-548c-5acd-9a6b-08d35b4ef87b")},
			{Sid: models.Sid("7a53aa91-bac9-5be6-9535-a2abf5fafc3b")},
			{Sid: models.Sid("77f1861e-faf0-56fa-aded-0519ce7d272c")},
			{Sid: models.Sid("23464ae4-d1db-579f-8625-97a52a83073a")},
			{Sid: models.Sid("4472673c-dd73-5ae7-ae69-22a25c16789e")},
			{Sid: models.Sid("2158d529-8fc2-57aa-adf4-f2001e6706ee")},
			{Sid: models.Sid("33d12640-8eea-5bba-afa0-496cdb79d210")},
			{Sid: models.Sid("40ed72db-3364-55cd-846c-f43bd8a2738a")},
			{Sid: models.Sid("5aef915a-2c61-528a-826b-ea95b3535788")},
			{Sid: models.Sid("d8d2225d-00d6-55ae-948c-efa4acdb6184")},
			{Sid: models.Sid("5f3c9078-6747-5551-a039-e0f73eda6dae")},
			{Sid: models.Sid("c0df50cd-ec6f-55f3-95a0-6a61d69569ef")},
			{Sid: models.Sid("c3f7eb98-ea8b-59b6-93de-36ec29c64ff0")},
			{Sid: models.Sid("817e5d63-d239-5415-ae0b-c089ad940ff7")},
			{Sid: models.Sid("4aa3f41e-6f17-536b-82fc-a2f21c1c0ec8")},
			{Sid: models.Sid("2eb31d60-3275-5d6f-8634-6ffd0594a256")},
			{Sid: models.Sid("2907a02c-8f1d-55ff-83fe-18eab82b1f1c")},
			{Sid: models.Sid("fe6c3d6c-6f7a-5a12-96cf-0a5966d8d229")},
			{Sid: models.Sid("de8fe33a-bc14-578b-b370-03a4e6ed115f")},
			{Sid: models.Sid("abb1fe0b-accd-5b2b-a0b8-c8e23d66d9b0")},
			{Sid: models.Sid("1c92476f-a3c5-59ae-a4d0-9539be1be356")},
			{Sid: models.Sid("9d1ff7c1-ebad-552d-a67c-b4670c9375e4")},
			{Sid: models.Sid("a9dca62e-a013-5a79-b5f7-2cbaa5d1a9db")},
			{Sid: models.Sid("5808b0a9-036e-5732-8666-d4865c28ea4f")},
			{Sid: models.Sid("8eac67e4-b989-5751-9c35-6d0d1fef7839")},
			{Sid: models.Sid("d148f8bc-941b-58ee-9774-7c3d245c2e67")},
			{Sid: models.Sid("345a6657-5b13-58a1-911e-7748a4b7c184")},
			{Sid: models.Sid("28cb4f02-b830-5d68-b9c4-b57c1d13f41b")},
			{Sid: models.Sid("369e4e06-9e3c-59a8-b794-fc37959429b3")},
			{Sid: models.Sid("2ce32209-dbd5-5288-bf83-982a52ab00fa")},
			{Sid: models.Sid("673a1204-0f38-5e92-b1fb-24189c043751")},
			{Sid: models.Sid("c009498b-102e-53cf-9508-88aec2e3848e")},
			{Sid: models.Sid("76420b57-f84c-5719-9c96-3e0e7197bde4")},
			{Sid: models.Sid("c77dc427-955d-5655-a1ca-48c88072311d")},
			{Sid: models.Sid("3cdb829f-26f2-572a-825f-aaa85d1112d0")},
			{Sid: models.Sid("7f17bedf-5398-5dbf-984b-5b1b99f5d44a")},
			{Sid: models.Sid("87652ba8-9f16-5721-8328-ff89edeca2bd")},
			{Sid: models.Sid("94658f68-44c5-53da-bea4-33f1e015f926")},
			{Sid: models.Sid("c2192ecb-9125-5906-8db6-aae97f65be70")},
		},
	}

	for noBatches := 2; noBatches < 10; noBatches++ {
		subBatches := batch.Divide(noBatches)

		if len(subBatches) != noBatches {
			t.Errorf("expected %d batches, got %d", noBatches, len(subBatches))
		}

		if err := checkAllItemsPlaced(batch, subBatches); err != nil {
			t.Error(errors.Wrap(err, fmt.Sprintf("noBatches: %d", noBatches)))
		}

		for _, subBatch := range subBatches {
			if err := checkBatchSize(batch, subBatch, noBatches); err != nil {
				t.Error(errors.Wrap(err, fmt.Sprintf("noBatches: %d", noBatches)))
			}
			if err := checkItems(batch, subBatch); err != nil {
				t.Error(errors.Wrap(err, fmt.Sprintf("noBatches: %d", noBatches)))
			}
		}

	}
}
