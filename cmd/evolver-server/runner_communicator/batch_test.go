package runner_communicator

import (
	"fmt"
	"tde/models"
	"testing"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

func Test_BatchDivide(t *testing.T) {
	checkBatchSize := func(batch *Batch, subBatch *Batch, noBatches int) error {
		maxExpected := len(batch.Candidates)/noBatches + 1
		minExpected := len(batch.Candidates) / noBatches

		if len(subBatch.Candidates) > maxExpected || len(subBatch.Candidates) < minExpected {
			return fmt.Errorf("batch size out of range, expected %d <= size <= %d, got %d", minExpected, maxExpected, len(subBatch.Candidates))
		}
		return nil
	}

	checkAllItemsPlaced := func(batch *Batch, subBatches []*Batch) error {
		merged := Batch{}
		for _, subBatch := range subBatches {
			merged.Candidates = append(merged.Candidates, subBatch.Candidates...)
		}

		for _, item := range batch.Candidates {
			if slices.Index(merged.Candidates, item) == -1 {
				return fmt.Errorf("None of the subbatches has the item: %v", item)
			}
		}

		return nil
	}

	checkItems := func(batch *Batch, subBatch *Batch) error {
		for _, c := range subBatch.Candidates {
			found := false
			for _, cc := range batch.Candidates {
				if c == cc {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("candidate %v not found in original batch", c)
			}
		}
		return nil
	}

	batch := &Batch{
		FileTemplate: "package foo, decls, stmts, exprs",
		Candidates: []*models.Candidate{
			{UUID: models.CandidateID("0a40f3cd-5d02-594c-bf1a-6c5aca05edb8")},
			{UUID: models.CandidateID("2901cbd4-72e7-5b97-9fa0-d50f863f2d0b")},
			{UUID: models.CandidateID("c33efc69-0130-5301-b242-1dbe83e53e42")},
			{UUID: models.CandidateID("31bc5f21-67f8-5088-bc52-4056295fe96a")},
			{UUID: models.CandidateID("5f074b1d-3b91-50cd-a11e-e8a5e8f1ac36")},
			{UUID: models.CandidateID("32817d78-2e71-57b8-8c3c-5ff704b17b42")},
			{UUID: models.CandidateID("c0439809-9087-50ed-a296-a490e615cde0")},
			{UUID: models.CandidateID("1efdb10e-8693-5992-9435-fd958bfa372d")},
			{UUID: models.CandidateID("9525d2fb-c448-5db3-9983-ecb9ac5345d7")},
			{UUID: models.CandidateID("3b1108f2-c4fe-5cde-a2d7-0360caebdd5a")},
			{UUID: models.CandidateID("b1e67897-b964-5440-8440-ce0af1185a5c")},
			{UUID: models.CandidateID("c3a26537-0159-525f-868c-9cb7b370b276")},
			{UUID: models.CandidateID("9c5b6e93-ab69-5ff2-bbde-d9f49e2dcd63")},
			{UUID: models.CandidateID("4a5024ec-a546-5c5b-9835-d191754df6b1")},
			{UUID: models.CandidateID("60ea077b-6789-531d-a967-c226bb7fa69f")},
			{UUID: models.CandidateID("e971e216-873b-5907-8d6b-76171c886405")},
			{UUID: models.CandidateID("f2da10ed-a825-55b2-ac52-0670aae41f56")},
			{UUID: models.CandidateID("805f6fe1-486d-5386-90bd-4ec492e0e2b9")},
			{UUID: models.CandidateID("d2481f8a-264a-5d26-832e-adf73dac328b")},
			{UUID: models.CandidateID("144e1c30-1ee3-577e-9c5a-bc1435ab6b06")},
			{UUID: models.CandidateID("e1b7e032-e13b-5991-916d-eec78fe02543")},
			{UUID: models.CandidateID("5aa14327-8395-509f-a500-18634f134bc7")},
			{UUID: models.CandidateID("adc3fa48-afa6-5171-9043-bb303664a2f3")},
			{UUID: models.CandidateID("8912df97-9120-5951-b2da-d3749ccdfccf")},
			{UUID: models.CandidateID("a5705dec-3529-5381-8d42-38fbcd364f61")},
			{UUID: models.CandidateID("eccd8712-618e-5144-8daf-732168ef9c31")},
			{UUID: models.CandidateID("7afd0398-d554-502f-ba6d-5fc0c415e5ea")},
			{UUID: models.CandidateID("9ae3d665-18bf-5206-97a5-ce6316e675a9")},
			{UUID: models.CandidateID("be60b933-1fe6-53cf-9c95-8237e87ac899")},
			{UUID: models.CandidateID("fec310c4-0fc0-51ba-b2bf-4990ca3b8695")},
			{UUID: models.CandidateID("baf80a78-f07e-5308-af59-27b62003d6ee")},
			{UUID: models.CandidateID("be051481-c079-598a-b28a-9fd2bd899494")},
			{UUID: models.CandidateID("7cbaade0-39c9-5298-bc11-521e3b947fc0")},
			{UUID: models.CandidateID("5fd9567e-04de-50ef-ba3f-a2194869b0ab")},
			{UUID: models.CandidateID("dc9b920d-7734-57a3-9183-0599f2aaf2b1")},
			{UUID: models.CandidateID("9a23f972-a3b4-56c5-87f6-e23d75ab8085")},
			{UUID: models.CandidateID("b534d245-4c24-57fb-ae51-2520d60e675d")},
			{UUID: models.CandidateID("56188b2e-834e-557f-819e-345ece3eaf7b")},
			{UUID: models.CandidateID("15317383-eb25-5001-99d8-0c93185c3138")},
			{UUID: models.CandidateID("a28b1f49-84de-5a15-b756-a1bf521e116b")},
			{UUID: models.CandidateID("4adf206c-d58d-5185-8808-1f9785985b98")},
			{UUID: models.CandidateID("ee24fb34-846e-57b6-9015-e915317b350e")},
			{UUID: models.CandidateID("f58272e2-37a1-5614-b52d-0c3b6ca0bc4c")},
			{UUID: models.CandidateID("a8ea22fd-ea3c-5e9e-8a6a-ea3df044467a")},
			{UUID: models.CandidateID("36e5accd-5d41-53b6-8df2-1cd8ef10dfc1")},
			{UUID: models.CandidateID("e06e593a-14b0-5070-9b2c-e0ceb073f4be")},
			{UUID: models.CandidateID("69e79178-0b02-5a2b-8a9f-10e4d1e5672e")},
			{UUID: models.CandidateID("13f27439-a1ee-5c40-879d-b07fa1466a38")},
			{UUID: models.CandidateID("b312f0e0-af9f-535c-91a2-89b9ae1bf5fa")},
			{UUID: models.CandidateID("883478ee-c1ef-51cd-a576-3daaf003fdb7")},
			{UUID: models.CandidateID("4b4c57dd-82d7-5054-a92a-d3ed643c0196")},
			{UUID: models.CandidateID("41b737e1-923a-5802-8544-449c7bcea851")},
			{UUID: models.CandidateID("c4dd644e-8fc7-5fd8-bd45-4165416896bc")},
			{UUID: models.CandidateID("b107c4e8-feba-5745-a5c9-14166b313915")},
			{UUID: models.CandidateID("56a50afc-57f1-5837-b2da-3563b5c59fbe")},
			{UUID: models.CandidateID("e75a719b-3966-52c0-82a4-0e728a353d1c")},
			{UUID: models.CandidateID("eef87cf7-acde-5fd1-b052-209f8e8e7348")},
			{UUID: models.CandidateID("d5881b1d-cd93-55b0-a788-eb7285883e45")},
			{UUID: models.CandidateID("8f6dc741-5a65-5d83-b0f5-feb21f72c432")},
			{UUID: models.CandidateID("f2ab924e-2fc5-55f9-b748-add728861457")},
			{UUID: models.CandidateID("b0098db4-b2cd-5922-b6cc-82b8137ae3e8")},
			{UUID: models.CandidateID("2d1bd358-540c-5ee6-81fb-f0bc4837e3ea")},
			{UUID: models.CandidateID("b66cc41a-40bc-5cc7-8fd7-328028906925")},
			{UUID: models.CandidateID("f68f5375-3205-5eb4-aa65-862f8290de69")},
			{UUID: models.CandidateID("d0ddd3c9-56c1-5db0-9196-916a92d18e62")},
			{UUID: models.CandidateID("5547d16c-fb50-519e-85ce-ce66ae815360")},
			{UUID: models.CandidateID("3beb9680-5cad-550c-9cb5-e98a4109a560")},
			{UUID: models.CandidateID("a4d2aea9-a629-5f7b-a721-6d317f95a59d")},
			{UUID: models.CandidateID("e9a53135-1a89-5d84-b717-7eff6343339b")},
			{UUID: models.CandidateID("9f19e3fa-22b2-5705-827d-0c2cd6ae37bb")},
			{UUID: models.CandidateID("10b6a83d-c9bb-57b1-92e7-ac40efd0a87f")},
			{UUID: models.CandidateID("e1b83767-c208-5fb5-8cf4-749f1070f9c7")},
			{UUID: models.CandidateID("7e4f9969-1462-59d8-8fac-035c47fab00e")},
			{UUID: models.CandidateID("4d155049-0601-5c01-a846-6e16391ada6b")},
			{UUID: models.CandidateID("28c2499a-306c-5945-9ed8-394fc98f019f")},
			{UUID: models.CandidateID("46a06cfa-db5f-5944-839a-91b34841c9e7")},
			{UUID: models.CandidateID("4ad6ee04-f28f-5872-82d5-8284bf5cc8f2")},
			{UUID: models.CandidateID("a2fe346f-5c84-51dd-a67d-16cc3fc2c6c7")},
			{UUID: models.CandidateID("b24a8a0a-262a-546f-97c5-5656c1fbb8f8")},
			{UUID: models.CandidateID("575e4ab7-597f-500b-b9d4-c5ffc2735047")},
			{UUID: models.CandidateID("ee87a6a5-a050-5483-964e-205caadd131d")},
			{UUID: models.CandidateID("bc59aeb5-c093-5a2b-8954-9917eda3e51d")},
			{UUID: models.CandidateID("c5ee728e-e040-571d-92c5-5347b87ecb42")},
			{UUID: models.CandidateID("c5938cd7-5a21-5f84-8e96-c2115f27645b")},
			{UUID: models.CandidateID("4bca3803-2e2a-5443-aaf8-3bae4db0fbf6")},
			{UUID: models.CandidateID("77522330-93c2-5f52-8f8b-ec224f9f7289")},
			{UUID: models.CandidateID("e66fc896-207e-552b-8f73-d49b3c19e367")},
			{UUID: models.CandidateID("c4954209-89a2-571b-ac94-5705cc7e1883")},
			{UUID: models.CandidateID("286bd405-792a-5c82-9627-e27c102113d0")},
			{UUID: models.CandidateID("6faea496-ba7a-5080-a284-a93251b43b7a")},
			{UUID: models.CandidateID("d74d5dfd-a7f4-5dd3-a267-d9dc7c59af2d")},
			{UUID: models.CandidateID("b55e3353-13d6-5070-a05f-6fea4334608b")},
			{UUID: models.CandidateID("d2f284e6-9947-5e8e-bc8d-4a1b1354afb3")},
			{UUID: models.CandidateID("f594d589-49b7-5f52-b249-3661bbfdd049")},
			{UUID: models.CandidateID("17cc9928-766f-582b-a4df-7f1dd8d192a1")},
			{UUID: models.CandidateID("5c6fe1b4-33d7-51d1-aae8-c5877518e776")},
			{UUID: models.CandidateID("0cf27cd5-f742-51f6-ba2c-63841df9361f")},
			{UUID: models.CandidateID("974b2b04-4631-5c73-acc4-3a7eba66ca99")},
			{UUID: models.CandidateID("d105d930-eb06-5a16-9620-0971ac10eb2e")},
			{UUID: models.CandidateID("15168a45-a46d-5ea1-994d-6207ee901adf")},
			{UUID: models.CandidateID("21325c48-867c-5d9e-bb22-5093c42981db")},
			{UUID: models.CandidateID("aab4697c-2995-5ec0-b2eb-bbf487a74c5f")},
			{UUID: models.CandidateID("b48dc5c0-c29e-5e93-b9a6-e26c1e11c7d3")},
			{UUID: models.CandidateID("8d5300ca-4e1e-59ce-9ee2-67c60965daa3")},
			{UUID: models.CandidateID("3d689854-5d1a-5379-afd0-45fbaf1f10a1")},
			{UUID: models.CandidateID("c002a522-0212-523b-bc1a-1549bdda069f")},
			{UUID: models.CandidateID("6dabf5ae-8c85-57e6-8d28-94acd0ddf403")},
			{UUID: models.CandidateID("87e9f329-6e54-5db0-bf44-2a609b1dfcbb")},
			{UUID: models.CandidateID("3a5e0aef-ed28-5048-8ce5-5557e4c3938b")},
			{UUID: models.CandidateID("50e06961-cd3d-518e-8a69-77da74e91f36")},
			{UUID: models.CandidateID("04a8cb6d-4ce7-5263-911c-8cc95ed4c5d0")},
			{UUID: models.CandidateID("bb15b989-eb81-52a7-8478-b885201465fb")},
			{UUID: models.CandidateID("022625af-331d-5fdf-87bb-5bafbc22bbe1")},
			{UUID: models.CandidateID("f04977e8-4a66-518d-bed7-9ece284a6c6f")},
			{UUID: models.CandidateID("350f6111-aa3d-5800-a4a9-7f00f5114e34")},
			{UUID: models.CandidateID("fb3b0f36-368c-5bc1-85b5-141414c9924b")},
			{UUID: models.CandidateID("6e693c79-93c9-5aa5-b8da-5ec75dbabf95")},
			{UUID: models.CandidateID("2e6c96e3-0b62-5d89-ac33-4517c004361c")},
			{UUID: models.CandidateID("c8f92e45-3a96-5376-8a7d-aa9de564a72d")},
			{UUID: models.CandidateID("a64915c9-8d96-500d-9bdf-d0649b0da33a")},
			{UUID: models.CandidateID("f972d8fe-0c6a-5fe5-a535-e0050ae6072d")},
			{UUID: models.CandidateID("dcea01a6-ea0c-5ab3-9f2f-324403a5ec77")},
			{UUID: models.CandidateID("a8dd4574-e90a-595e-a71f-84fbfb561712")},
			{UUID: models.CandidateID("b83451bb-494a-5a4c-b888-4d4175d5c37a")},
			{UUID: models.CandidateID("72c8d346-186a-52e6-976a-4885fdc09cc8")},
			{UUID: models.CandidateID("90dcb845-9740-5e29-8f13-76e48884cb0a")},
			{UUID: models.CandidateID("a2ffee5a-4948-5104-9a44-d76d12010a66")},
			{UUID: models.CandidateID("679df4cb-dba8-58df-a234-ba7ab5f68512")},
			{UUID: models.CandidateID("2f968ef3-63a0-5865-a8a4-144fe97de00f")},
			{UUID: models.CandidateID("c0bb3325-d38e-5126-822e-bb4339645e08")},
			{UUID: models.CandidateID("c273fd73-d3c7-5f0d-a4b8-5b2b9a94d133")},
			{UUID: models.CandidateID("73ad205f-5f20-5cf6-b0b8-4af69998ef8e")},
			{UUID: models.CandidateID("f68f3efd-aac8-5c6f-94f0-f32549dba2bd")},
			{UUID: models.CandidateID("0a46e858-5500-5342-b5a1-251d8c8bec47")},
			{UUID: models.CandidateID("f42058fc-513e-56db-821d-738ac8695e81")},
			{UUID: models.CandidateID("d8992870-9cb7-55e0-9794-2aa857c08efb")},
			{UUID: models.CandidateID("d362b622-3ba8-5a35-bcdd-6aed1eff0f8b")},
			{UUID: models.CandidateID("99d322ae-cf03-5011-8324-fb7f7faa4eb6")},
			{UUID: models.CandidateID("6b2843d6-7631-5928-a656-8f9aae6c27dc")},
			{UUID: models.CandidateID("98158575-0013-5b4e-a41a-b216d97ac71c")},
			{UUID: models.CandidateID("ae2d6ca1-71ae-562f-998d-0194b665d1ab")},
			{UUID: models.CandidateID("a0112f2b-fb7c-52d4-bd39-4a0ec09fedfc")},
			{UUID: models.CandidateID("a48df952-c834-5969-b093-dfd7dc51f033")},
			{UUID: models.CandidateID("a2519c4d-fe3d-5ec3-ade3-5e55265c0a51")},
			{UUID: models.CandidateID("74ac5484-ff48-5d7b-b1c1-b5f681c6687c")},
			{UUID: models.CandidateID("a02ca99d-bb45-582d-a132-0224e02e4b73")},
			{UUID: models.CandidateID("27e0152c-4990-5b1a-ac1a-fa0bb3143b4d")},
			{UUID: models.CandidateID("1719c34b-de61-59fb-8205-aa36e45a6f1c")},
			{UUID: models.CandidateID("0a987463-eafe-5302-ba70-b160c0b96b96")},
			{UUID: models.CandidateID("7deae393-bc39-524d-a331-98ada76b19cb")},
			{UUID: models.CandidateID("3ece7013-10e3-5f45-8a0a-25e71da6c2da")},
			{UUID: models.CandidateID("f165878f-ca32-53d8-8cbb-aa43af8d8c3c")},
			{UUID: models.CandidateID("d6235aa3-d699-50dd-9f2f-2159c99679cd")},
			{UUID: models.CandidateID("ae98b610-c211-5533-820d-4c764a5aa16f")},
			{UUID: models.CandidateID("efd562cf-2502-5e2c-8888-0b1cd6519736")},
			{UUID: models.CandidateID("f69eb48d-e864-5077-acde-56d25300e796")},
			{UUID: models.CandidateID("9f02b12b-6300-57a1-b3c8-7c0d5271f21b")},
			{UUID: models.CandidateID("aad64c35-57b9-5e92-985f-aee305a2e4cf")},
			{UUID: models.CandidateID("ee2f039e-c227-5672-86a3-848eadc24989")},
			{UUID: models.CandidateID("f4b621f8-ea54-51be-9abe-6cb58b09d3e1")},
			{UUID: models.CandidateID("4838efba-81cc-5728-b58c-af08f24f1659")},
			{UUID: models.CandidateID("95df06ca-c4ea-5402-9f29-1464751f8acd")},
			{UUID: models.CandidateID("3570279c-5f32-546e-8049-5b2f5f61d154")},
			{UUID: models.CandidateID("512abd39-c8f1-5fce-a1c0-242290da5db6")},
			{UUID: models.CandidateID("064ae3c0-2bde-5d01-adc5-8f84d0870b3f")},
			{UUID: models.CandidateID("3ef087c5-8ee7-5004-8467-1e156c736cfc")},
			{UUID: models.CandidateID("c7d0a1bf-08c0-56df-94ee-a946de92cc89")},
			{UUID: models.CandidateID("640444ae-941a-5f9b-b1bf-080c209a872f")},
			{UUID: models.CandidateID("9a02dc00-1b8d-572f-9eba-fc1ae299fb3a")},
			{UUID: models.CandidateID("88029ccc-9847-519f-ba80-2d95279cc6ca")},
			{UUID: models.CandidateID("0249864e-e700-520a-89d7-dc14b5d0dc9c")},
			{UUID: models.CandidateID("f7eb1d3a-d323-588e-864f-08b558f387b9")},
			{UUID: models.CandidateID("c0dab5bb-08ed-5027-887a-78deabdbec4f")},
			{UUID: models.CandidateID("a5520932-c935-57ed-93e5-7c4407f4f7b7")},
			{UUID: models.CandidateID("0627cd46-fb12-5105-9272-33a86c3ff617")},
			{UUID: models.CandidateID("00c07fac-2c54-5d8c-ab95-006fd57f3db7")},
			{UUID: models.CandidateID("4c441002-1cf7-58e6-8173-9d77e55644df")},
			{UUID: models.CandidateID("a3276500-12df-58e3-af27-2be4defcd700")},
			{UUID: models.CandidateID("7593be2f-a788-590d-b394-8d0ba37e3bd0")},
			{UUID: models.CandidateID("e7f56405-991f-5a72-934c-b3bb1009327a")},
			{UUID: models.CandidateID("5531222d-dd90-58e0-940f-dae664749957")},
			{UUID: models.CandidateID("49fa3495-900c-50fd-89b5-6dfdbfba2fc6")},
			{UUID: models.CandidateID("65f6f72a-02a6-50bd-9f5f-e50fae823fc8")},
			{UUID: models.CandidateID("9dea40cf-9031-5c31-ae26-debac8449e1d")},
			{UUID: models.CandidateID("dd69c590-8a64-54e6-ba55-e33b159a92f9")},
			{UUID: models.CandidateID("495eecc5-36db-5066-a95a-771e0654eb37")},
			{UUID: models.CandidateID("2771ff78-ebad-5bd6-a1e8-df42f417d22e")},
			{UUID: models.CandidateID("22a9fa04-8c9a-5d20-89d1-8248c42d686e")},
			{UUID: models.CandidateID("2e5052f6-1ee0-5b45-b533-2303ee30e05d")},
			{UUID: models.CandidateID("0b1ba8d8-f0be-521b-b845-216aa48eb871")},
			{UUID: models.CandidateID("81a62c4a-2061-5fb9-83b9-bd75ac77f30f")},
			{UUID: models.CandidateID("f5102884-1117-5059-99d1-1a1f1461b062")},
			{UUID: models.CandidateID("bbb12c3e-395e-5233-8720-cfd9161f6ccb")},
			{UUID: models.CandidateID("6f9e562b-e394-59f0-bac4-26b0953c1653")},
			{UUID: models.CandidateID("97f03b6f-d842-578c-ad2d-2b65ed6645d4")},
			{UUID: models.CandidateID("b99ae5c2-d232-53eb-ba66-925ecee5e3ea")},
			{UUID: models.CandidateID("ba394f36-c462-578d-998f-3bd05df73275")},
			{UUID: models.CandidateID("e2059873-dbee-5009-97c0-8ed9406ebd07")},
			{UUID: models.CandidateID("df73b8a6-522a-5cdc-9fce-cf1d9fff761e")},
			{UUID: models.CandidateID("35301f0b-a0ca-5259-916f-7fbf862cce29")},
			{UUID: models.CandidateID("affbf134-25f6-5b26-8c05-90b757713a8f")},
			{UUID: models.CandidateID("6e558925-2fc0-56c3-8e0c-811e08090e57")},
			{UUID: models.CandidateID("a7edf474-e5dd-5995-a925-be41e2fef200")},
			{UUID: models.CandidateID("3292f169-e8cd-595d-9fac-91d3078f1f15")},
			{UUID: models.CandidateID("b7720e1c-cfeb-51cb-9ab4-bd10dfe9f76c")},
			{UUID: models.CandidateID("f7771b88-dd1a-5701-8313-de13fc57700b")},
			{UUID: models.CandidateID("9c294e5b-b4d6-51f7-937f-bab6fc0ebb86")},
			{UUID: models.CandidateID("9add59fc-c9c2-598f-83e7-d81a0738fe29")},
			{UUID: models.CandidateID("c6e62b39-b9f3-5158-8bb5-e3ffb30c7ae8")},
			{UUID: models.CandidateID("76820be0-9625-5022-a77d-e7e934277b13")},
			{UUID: models.CandidateID("fe028060-7b49-51eb-9e63-79e8825a47f8")},
			{UUID: models.CandidateID("36b05b29-bb04-546e-8e59-4c67f40046c4")},
			{UUID: models.CandidateID("35e18602-412f-5f8f-8921-91d426d588e1")},
			{UUID: models.CandidateID("26ce239b-e4e0-5220-a1e3-68aa7b0545fd")},
			{UUID: models.CandidateID("53233f72-7fa7-5b9b-a569-009f54574898")},
			{UUID: models.CandidateID("52e91c09-a26c-5604-9ede-5a1aba029106")},
			{UUID: models.CandidateID("0fb13e54-5837-5425-882e-1eb844e4c7cc")},
			{UUID: models.CandidateID("a30e20a1-f1f3-55d8-bbc5-7ded4bf3d6c7")},
			{UUID: models.CandidateID("8fdaf5f7-95f4-53cc-a11a-ea2132bdfa6d")},
			{UUID: models.CandidateID("a61ff288-02d4-518f-9a41-5bcba7afb957")},
			{UUID: models.CandidateID("163d3ef4-b087-5c1d-b5d4-d4051357150f")},
			{UUID: models.CandidateID("7d5ff472-9464-5e65-9f52-9cf28907277e")},
			{UUID: models.CandidateID("c9b69e3a-601e-5398-93a9-169b453b47ee")},
			{UUID: models.CandidateID("219b6c8c-73d1-5971-a77c-dd0e0ff39df6")},
			{UUID: models.CandidateID("3c35ef39-aa72-5374-a291-5099bf23e4ee")},
			{UUID: models.CandidateID("0d9c7701-cf73-572b-9a7c-ef35d8ab472c")},
			{UUID: models.CandidateID("d712c83f-77d6-504b-b0cf-c5e0d5656b17")},
			{UUID: models.CandidateID("c948e183-861c-5d73-bddc-d4308b53fc55")},
			{UUID: models.CandidateID("0e92952a-381f-53b3-9755-ca1b428b17a0")},
			{UUID: models.CandidateID("97d6e52a-56d0-5698-b6b1-ffba6e74722e")},
			{UUID: models.CandidateID("8056df46-6305-504f-a19c-51702020c92c")},
			{UUID: models.CandidateID("c981a644-5805-576c-ab58-481745848e09")},
			{UUID: models.CandidateID("d38115cd-9596-5961-beae-19805dc1406d")},
			{UUID: models.CandidateID("a320373e-eb21-5c10-bc56-0798b835907c")},
			{UUID: models.CandidateID("fd12420a-ce36-5f24-bb3f-7266620d250b")},
			{UUID: models.CandidateID("6762145b-dccc-50eb-8ea0-22e2bf9e11d7")},
			{UUID: models.CandidateID("25dd84c4-1f29-5189-b874-4508e1bd632a")},
			{UUID: models.CandidateID("6a9fd923-8b2f-5f5e-b888-e67fd0816d72")},
			{UUID: models.CandidateID("06cdba9c-e180-53e3-b691-e0b5f0155f01")},
			{UUID: models.CandidateID("98552a95-9745-5b20-8a38-41619eee8600")},
			{UUID: models.CandidateID("23f9b648-fa5f-5b7e-a0e9-8b1d2257efd8")},
			{UUID: models.CandidateID("cca4886e-c5db-5db8-a56f-6cd9cd311447")},
			{UUID: models.CandidateID("50cbabdd-0093-54e4-92bc-441402fe5cc8")},
			{UUID: models.CandidateID("589c922a-aa71-57da-b59e-d1e484f09a3d")},
			{UUID: models.CandidateID("6f13f7bf-d1bf-5e0e-9c13-949bf0f93958")},
			{UUID: models.CandidateID("20d75553-321e-5386-bbc6-12267f761b22")},
			{UUID: models.CandidateID("365c9b20-39c1-5c5e-a39c-7381ba65b7b2")},
			{UUID: models.CandidateID("eb7f96da-548c-5acd-9a6b-08d35b4ef87b")},
			{UUID: models.CandidateID("7a53aa91-bac9-5be6-9535-a2abf5fafc3b")},
			{UUID: models.CandidateID("77f1861e-faf0-56fa-aded-0519ce7d272c")},
			{UUID: models.CandidateID("23464ae4-d1db-579f-8625-97a52a83073a")},
			{UUID: models.CandidateID("4472673c-dd73-5ae7-ae69-22a25c16789e")},
			{UUID: models.CandidateID("2158d529-8fc2-57aa-adf4-f2001e6706ee")},
			{UUID: models.CandidateID("33d12640-8eea-5bba-afa0-496cdb79d210")},
			{UUID: models.CandidateID("40ed72db-3364-55cd-846c-f43bd8a2738a")},
			{UUID: models.CandidateID("5aef915a-2c61-528a-826b-ea95b3535788")},
			{UUID: models.CandidateID("d8d2225d-00d6-55ae-948c-efa4acdb6184")},
			{UUID: models.CandidateID("5f3c9078-6747-5551-a039-e0f73eda6dae")},
			{UUID: models.CandidateID("c0df50cd-ec6f-55f3-95a0-6a61d69569ef")},
			{UUID: models.CandidateID("c3f7eb98-ea8b-59b6-93de-36ec29c64ff0")},
			{UUID: models.CandidateID("817e5d63-d239-5415-ae0b-c089ad940ff7")},
			{UUID: models.CandidateID("4aa3f41e-6f17-536b-82fc-a2f21c1c0ec8")},
			{UUID: models.CandidateID("2eb31d60-3275-5d6f-8634-6ffd0594a256")},
			{UUID: models.CandidateID("2907a02c-8f1d-55ff-83fe-18eab82b1f1c")},
			{UUID: models.CandidateID("fe6c3d6c-6f7a-5a12-96cf-0a5966d8d229")},
			{UUID: models.CandidateID("de8fe33a-bc14-578b-b370-03a4e6ed115f")},
			{UUID: models.CandidateID("abb1fe0b-accd-5b2b-a0b8-c8e23d66d9b0")},
			{UUID: models.CandidateID("1c92476f-a3c5-59ae-a4d0-9539be1be356")},
			{UUID: models.CandidateID("9d1ff7c1-ebad-552d-a67c-b4670c9375e4")},
			{UUID: models.CandidateID("a9dca62e-a013-5a79-b5f7-2cbaa5d1a9db")},
			{UUID: models.CandidateID("5808b0a9-036e-5732-8666-d4865c28ea4f")},
			{UUID: models.CandidateID("8eac67e4-b989-5751-9c35-6d0d1fef7839")},
			{UUID: models.CandidateID("d148f8bc-941b-58ee-9774-7c3d245c2e67")},
			{UUID: models.CandidateID("345a6657-5b13-58a1-911e-7748a4b7c184")},
			{UUID: models.CandidateID("28cb4f02-b830-5d68-b9c4-b57c1d13f41b")},
			{UUID: models.CandidateID("369e4e06-9e3c-59a8-b794-fc37959429b3")},
			{UUID: models.CandidateID("2ce32209-dbd5-5288-bf83-982a52ab00fa")},
			{UUID: models.CandidateID("673a1204-0f38-5e92-b1fb-24189c043751")},
			{UUID: models.CandidateID("c009498b-102e-53cf-9508-88aec2e3848e")},
			{UUID: models.CandidateID("76420b57-f84c-5719-9c96-3e0e7197bde4")},
			{UUID: models.CandidateID("c77dc427-955d-5655-a1ca-48c88072311d")},
			{UUID: models.CandidateID("3cdb829f-26f2-572a-825f-aaa85d1112d0")},
			{UUID: models.CandidateID("7f17bedf-5398-5dbf-984b-5b1b99f5d44a")},
			{UUID: models.CandidateID("87652ba8-9f16-5721-8328-ff89edeca2bd")},
			{UUID: models.CandidateID("94658f68-44c5-53da-bea4-33f1e015f926")},
			{UUID: models.CandidateID("c2192ecb-9125-5906-8db6-aae97f65be70")},
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
